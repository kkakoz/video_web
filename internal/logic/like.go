package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sync"
	"video_web/internal/async/producer"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
	"video_web/pkg/errno"
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

func (item *likeLogic) Like(ctx context.Context, req *dto.Like) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}

	err = ormx.Transaction(ctx, func(ctx context.Context) (err error) {

		defer func() {
			if err != nil {
				return
			}
			// 点赞成功发送kafka消息队列
			if req.LikeType && req.Like {
				err = producer.Send(&dto.Event{
					EventType:     dto.EventTypeLike,
					TargetId:      req.TargetId,
					TargetType:    req.TargetType,
					ActorId:       user.ID,
					TargetOwnerId: 0,
				})
				// 点赞成功发送重新计算hot
				_ = Video().AddHots(ctx, req.TargetId)
			}

		}()

		userLike, err := repo.Like().Get(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, req.TargetType, req.TargetId))
		if err != nil {
			return err
		}
		if req.LikeType { // 添加like / dislike
			if userLike == nil { // 第一次点赞
				err = repo.Like().Add(ctx, &entity.Like{
					UserId:     user.ID,
					TargetType: req.TargetType,
					TargetId:   req.TargetId,
					Like:       req.Like,
				})
				if err != nil {
					return err
				}
				if req.Like {
					return item.UpdateCount(ctx, req.TargetId, req.TargetType, 1)
				}
				return nil
			}

			//  已经点过赞
			if userLike.Like == req.Like { // 已经点过
				return errno.New400("请不要重复操作")
			} else {
				err = repo.Like().Updates(ctx, map[string]any{
					"like": req.Like,
				}, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, req.TargetType, req.TargetId))
				if err != nil {
					return err
				}
				updateCount := lo.Ternary(req.Like, 1, -1)
				err = item.UpdateCount(ctx, req.TargetId, req.TargetType, updateCount)
				return err
			}

		} else {
			if userLike != nil {
				err = repo.Like().Delete(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, req.TargetType, req.TargetId))
				if err != nil {
					return err
				}
				if userLike.Like {
					err = item.UpdateCount(ctx, req.TargetId, req.TargetType, -1)
					return err
				}
			}
			return nil
		}
	})
	if err != nil {
		return err
	}

	return nil
}

func (likeLogic) UpdateCount(ctx context.Context, id int64, targetType entity.TargetType, updateCount int) error {
	switch targetType {
	case entity.TargetTypeVideo:
		err := repo.Video().Updates(ctx, map[string]any{"like": gorm.Expr("`like` + ?", updateCount)},
			opt.Where("id = ?", id))
		return err
	}
	return nil
}

func (likeLogic) Likes(ctx context.Context, req *dto.LikeIs) ([]*entity.Like, error) {
	return repo.Like().GetList(ctx, opt.Where("user_id = ? and target_type = ? and target_id in ?", req.UserId, req.TargetType, req.TargetIds))
}
