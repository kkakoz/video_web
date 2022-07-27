package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sync"
	"video_web/internal/consts"
	"video_web/internal/dto/request"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model"
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

func (item *likeLogic) Like(ctx context.Context, req *request.LikeReq) error {
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		user, err := local.GetUser(ctx)
		if err != nil {
			return err
		}
		if req.LikeType { // 添加点赞
			err = repo.Like().Add(ctx, &model.Like{
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
		updateCount := lo.Ternary(req.LikeType, 1, -1)
		switch req.TargetType {
		case consts.LikeTypeVideo:
			err = repo.Collection().Updates(ctx, map[string]any{"like_count": gorm.Expr("like_count + ?", updateCount)},
				opt.Where("id = ?"))
		}
		return err
	})
}

func (item *likeLogic) IsLike(ctx context.Context, req *request.LikeIsReq) (bool, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return false, err
	}

	return repo.Like().GetExist(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, req.TargetType, req.TargetId))
}
