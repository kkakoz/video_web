package logic

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opts"
	"video_web/internal/domain"
	"video_web/internal/dto/request"
	"video_web/internal/pkg/keys"
	"video_web/pkg/local"
	"video_web/pkg/redis/bloom_filter"
)

var _ domain.ILikeLogic = (*LikeLogic)(nil)

type LikeLogic struct {
	likeRepo domain.ILikeRepo
	cache    *redis.Client
}

func NewLikeLogic(likeRepo domain.ILikeRepo, cache *redis.Client) domain.ILikeLogic {
	return &LikeLogic{likeRepo: likeRepo, cache: cache}
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
	if req.LikeType { // 加入 布尔过滤器
		err = item.likeRepo.Add(ctx, &domain.Like{
			UserId:     user.ID,
			TargetType: req.TargetType,
			TargetId:   req.TargetId,
		})
		if err != nil {
			return err
		}
		filter := bloom_filter.NewBloomFilter(item.cache)
		return filter.Add(keys.LikeValueKey(req.TargetType, req.TargetId), fmt.Sprintf("%d", user.ID))
	}
	return item.likeRepo.Delete(ctx, opts.Where("target_type = ? and target_id = ?", req.TargetType, req.TargetId))
}

func (item *LikeLogic) IsLike(ctx context.Context, req *request.LikeIsReq) (bool, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return false, err
	}
	filter := bloom_filter.NewBloomFilter(item.cache)
	b := filter.Contains(keys.LikeValueKey(req.TargetType, req.TargetId), fmt.Sprintf("%d", user.ID))
	if b {
		return item.likeRepo.GetExist(ctx, opts.Where("target_type = ? and target_id = ?", req.TargetType, req.TargetId))
	}
	return b, nil
}
