package repo

import (
	"context"
	"github.com/kkakoz/ormx"
	"gorm.io/gorm/clause"
	"sync"
	"video_web/internal/model/entity"
)

type likeRepo struct {
	ormx.IRepo[entity.Like]
}

var likeOnce sync.Once
var _like *likeRepo

func Like() *likeRepo {
	likeOnce.Do(func() {
		_like = &likeRepo{IRepo: ormx.NewRepo[entity.Like]()}
	})
	return _like
}

func (likeRepo) AddLikeState(ctx context.Context, list []*entity.Like) error {
	if len(list) == 0 {
		return nil
	}
	db := ormx.DB(ctx)
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "target_id"}, {Name: "target_type"}},
		DoUpdates: clause.AssignmentColumns([]string{"like", "like"}),
	}).Create(&list).Error
}

func (l likeRepo) DeleteCancelLikes(ctx context.Context, list []*entity.Like) error {
	if len(list) == 0 {
		return nil
	}
	db := ormx.DB(ctx)
	params := make([][]any, 0)
	for _, v := range list {
		params = append(params, []any{v.UserId, v.TargetId, v.TargetType})
	}

	return db.Where("(user_id, target_id, target_type) in ?", params).Delete(&entity.Like{}).Error
}
