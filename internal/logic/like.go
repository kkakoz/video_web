package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
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

func (likeLogic) Like(ctx context.Context, req *dto.Like) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}

	return ormx.Transaction(ctx, func(ctx context.Context) error {
		if req.LikeType { // 添加点赞
			err = repo.Like().Add(ctx, &entity.Like{
				UserId:     user.ID,
				TargetType: req.TargetType,
				TargetId:   req.TargetId,
			})
			return err
		} else {
			err = repo.Like().Delete(ctx, opt.Where("target_type = ? and target_id = ?", req.TargetType, req.TargetId))
			if err != nil {
				return err
			}
		}

		// 点赞数量更新
		updateCount := lo.Ternary(req.LikeType, 1, -1)
		switch req.TargetType {
		case entity.LikeTargetTypeVideo:
			err = repo.Resource().Updates(ctx, map[string]any{"like_count": gorm.Expr("like_count + ?", updateCount)},
				opt.Where("id = ?"))
		}
		return err
	})
}

func (likeLogic) Likes(ctx context.Context, req *dto.LikeIs) ([]*entity.Like, error) {
	return repo.Like().GetList(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", req.UserId, req.TargetType, req.TargetIds))
}
