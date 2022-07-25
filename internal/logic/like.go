package logic

import (
	"context"
	"fmt"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sync"
	"video_web/internal/consts"
	"video_web/internal/dto/request"
	repo2 "video_web/internal/logic/internal/repo"
	"video_web/internal/model"
	"video_web/internal/pkg/keys"
	"video_web/internal/pkg/local"
	"video_web/pkg/redisx"
	"video_web/pkg/redisx/bloom_filter"
)

type likeLogic struct {
}

var likeOnce sync.Once
var _like *likeLogic

func Like() *likeLogic {
	likeOnce.Do(func() {
		_like = &likeLogic{}
	})
	return _like
}

func (item *likeLogic) Like(ctx context.Context, req *request.LikeReq) (err error) {
	ctx, checkError := ormx.Begin(ctx)
	defer func() {
		err = checkError(err)
	}()
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	if req.LikeType { // 加入 布隆过滤器
		err = repo2.Like().Add(ctx, &model.Like{
			UserId:     user.ID,
			TargetType: req.TargetType,
			TargetId:   req.TargetId,
		})
		if err != nil {
			return err
		}
		filter := bloom_filter.NewBloomFilter(redisx.Client())
		return filter.Add(keys.LikeValueKey(req.TargetType, req.TargetId), fmt.Sprintf("%d", user.ID))
	} else {
		err = repo2.Like().Delete(ctx, opt.Where("target_type = ? and target_id = ?", req.TargetType, req.TargetId))
		if err != nil {
			return err
		}
	}
	updateCount := lo.Ternary(req.LikeType, 1, -1)
	switch req.TargetType {
	case consts.LikeTypeVideo:
		err = repo2.Video().Updates(ctx, map[string]any{"like_count": gorm.Expr("like_count + ?", updateCount)},
			opt.Where("id = ?"))
	}
	return
}

func (item *likeLogic) IsLike(ctx context.Context, req *request.LikeIsReq) (bool, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return false, err
	}

	filter := bloom_filter.NewBloomFilter(redisx.Client())
	b := filter.Contains(keys.LikeValueKey(req.TargetType, req.TargetId), fmt.Sprintf("%d", user.ID))
	if b { // 可能存在
		return repo2.Like().GetExist(ctx, opt.Where("target_type = ? and target_id = ?", req.TargetType, req.TargetId))
	}
	return b, nil
}
