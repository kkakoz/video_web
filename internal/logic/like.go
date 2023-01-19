package logic

import (
	"context"
	"fmt"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/kkakoz/pkg/logger"
	"github.com/kkakoz/pkg/redisx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"sync"
	"video_web/internal/async/producer"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/keys"
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
			if req.LikeState == 1 {
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

		userLike, err := item.IsLike(ctx, req.TargetType, req.TargetId, user.ID)
		if err != nil {
			return err
		}
		if userLike != nil && userLike.Like == req.LikeState {
			return errno.New400("请不要重复操作")
		}
		err = item.AddLikeToCache(ctx, req.TargetType, req.TargetId, user.ID, req.LikeState)
		if err != nil {
			return err
		}
		return nil
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

func (likeLogic) Likes(ctx context.Context, req *dto.UserLikeList) ([]*entity.Like, error) {
	return repo.Like().GetList(ctx, opt.Where("user_id = ? and target_type = ? and target_id in ?", req.UserId, req.TargetType, req.TargetIds))
}

func (item *likeLogic) IsLike(ctx context.Context, targetType entity.TargetType, targetId int64, userId int64) (*entity.Like, error) {
	likeResult, err := redisx.Client().HGet(keys.LikeHashCache(targetType, targetId), strconv.Itoa(int(userId))).Int()
	if err != nil {
		like, err := repo.Like().Get(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", userId, targetType, targetId))
		if err != nil {
			return nil, err
		}
		return like, nil
	}
	return &entity.Like{
		UserId:     userId,
		TargetType: targetType,
		TargetId:   targetId,
		Like:       entity.LikeState(likeResult),
	}, nil
}

func (item *likeLogic) AddLikeToCache(ctx context.Context, targetType entity.TargetType, targetId int64, userId int64, like entity.LikeState) error {
	pipeliner := redisx.Client().TxPipeline()
	pipeliner.HSet(keys.LikeHashCache(targetType, targetId), strconv.Itoa(int(userId)), fmt.Sprintf("%d", like))
	pipeliner.HSet(keys.LikeHashSync(targetType), fmt.Sprintf("%d:%d", targetId, userId), fmt.Sprintf("%d", like))
	_, err := pipeliner.Exec()
	return err
}

func (item *likeLogic) UpdateLikeJob(ctx context.Context) error {
	kv, err := redisx.Client().HGetAll(keys.LikeHashSync(entity.TargetTypeVideo)).Result()
	if err != nil {
		return err
	}

	videoMap := map[int][]*entity.Like{}
	deletes := make([]*entity.Like, 0)
	adds := make([]*entity.Like, 0)

	for k, v := range kv {
		split := strings.Split(k, ":")
		if len(split) != 2 {
			continue
		}
		targetId, err := strconv.Atoi(split[0])
		if err != nil {
			continue
		}
		userId, err := strconv.Atoi(split[1])
		if err != nil {
			continue
		}
		likeState, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		videoMap[targetId] = append(videoMap[targetId], &entity.Like{
			UserId:     int64(userId),
			TargetType: entity.TargetTypeVideo,
			TargetId:   int64(targetId),
			Like:       entity.LikeState(likeState),
		})
		if likeState == 0 { // cancel
			deletes = append(deletes, &entity.Like{
				UserId:     int64(userId),
				TargetType: entity.TargetTypeVideo,
				TargetId:   int64(targetId),
				Like:       entity.LikeState(likeState),
			})
		} else {
			adds = append(adds, &entity.Like{
				UserId:     int64(userId),
				TargetType: entity.TargetTypeVideo,
				TargetId:   int64(targetId),
				Like:       entity.LikeState(likeState),
			})
		}
	}

	for videoId, list := range videoMap {
		updateCount := 0
		for _, like := range list {
			if like.Like == entity.LikeStateLike {
				updateCount++
			} else if like.Like == entity.LikeStateCancel {
				updateCount--
			}
		}
		err := item.UpdateCount(ctx, int64(videoId), entity.TargetTypeVideo, updateCount)
		if err != nil {
			logger.Error("update video like count err:", zap.Error(err))
			continue
		}
	}

	err = repo.Like().AddLikeState(ctx, adds)
	if err != nil {
		logger.Error("add redis like data to db err:", zap.Error(err))
	}

	err = repo.Like().DeleteCancelLikes(ctx, deletes)
	if err != nil {
		logger.Error("delete redis like data to db err:", zap.Error(err))
	}

	return nil
}

//func (likeLogic) IsLike(ctx context.Context, req *dto.IsLike) bool {
//	return redisx.Client().HGet(keys.UserLikeKey())
//}
