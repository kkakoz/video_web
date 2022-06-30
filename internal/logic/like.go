package logic

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"video_web/internal/consts"
	"video_web/internal/dto/request"
	"video_web/internal/model"
	"video_web/internal/pkg/keys"
	"video_web/internal/repo"
	"video_web/pkg/local"
	"video_web/pkg/redis/bloom_filter"
)

type LikeLogic struct {
	likeRepo  *repo.LikeRepo
	cache     *redis.Client
	userRepo  *repo.UserRepo
	videoRepo *repo.VideoRepo
}

func NewLikeLogic(likeRepo *repo.LikeRepo, cache *redis.Client, userRepo *repo.UserRepo, videoRepo *repo.VideoRepo) *LikeLogic {
	return &LikeLogic{likeRepo: likeRepo, cache: cache, userRepo: userRepo, videoRepo: videoRepo}
}

func (item *LikeLogic) Like(ctx context.Context, req *request.LikeReq) (err error) {
	ctx, checkError := ormx.Begin(ctx)
	defer func() {
		err = checkError(err)
	}()
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	if req.LikeType { // 加入 布隆过滤器
		err = item.likeRepo.Add(ctx, &model.Like{
			UserId:     user.ID,
			TargetType: req.TargetType,
			TargetId:   req.TargetId,
		})
		if err != nil {
			return err
		}
		filter := bloom_filter.NewBloomFilter(item.cache)
		return filter.Add(keys.LikeValueKey(req.TargetType, req.TargetId), fmt.Sprintf("%d", user.ID))
	} else {
		err = item.likeRepo.Delete(ctx, opt.Where("target_type = ? and target_id = ?", req.TargetType, req.TargetId))
		if err != nil {
			return err
		}
	}
	updateCount := lo.Ternary(req.LikeType, 1, -1)
	switch req.TargetType {
	case consts.LikeTypeVideo:
		err = item.videoRepo.Updates(ctx, map[string]any{"like_count": gorm.Expr("like_count + ?", updateCount)},
			opt.Where("id = ?"))
	}
	return
}

func (item *LikeLogic) IsLike(ctx context.Context, req *request.LikeIsReq) (bool, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return false, err
	}

	filter := bloom_filter.NewBloomFilter(item.cache)
	b := filter.Contains(keys.LikeValueKey(req.TargetType, req.TargetId), fmt.Sprintf("%d", user.ID))
	if b { // 可能存在
		return item.likeRepo.GetExist(ctx, opt.Where("target_type = ? and target_id = ?", req.TargetType, req.TargetId))
	}
	return b, nil
}
