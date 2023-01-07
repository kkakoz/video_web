package repo

import (
	"context"
	"gorm.io/gorm/clause"
	"sync"
	"video_web/internal/model/entity"

	"github.com/kkakoz/ormx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type videoRepo struct {
	ormx.IRepo[entity.Video]
}

var videoOnce sync.Once
var _collection *videoRepo

func Video() *videoRepo {
	videoOnce.Do(func() {
		_collection = &videoRepo{IRepo: ormx.NewRepo[entity.Video]()}
	})
	return _collection
}

func (v *videoRepo) UpdateAfterOrderEpisode(ctx context.Context, videoId int64, index int64, updateVal int) error {
	db := ormx.DB(ctx)
	err := db.Model(&entity.Resource{}).Where("video_id = ? and index >= ?", videoId, index).
		Update("index", gorm.Expr("index + ?", updateVal)).Error
	return errors.Wrap(err, "更新失败")
}

func (v *videoRepo) UpdateHots(ctx context.Context, videos []*entity.Video) error {
	db := ormx.DB(ctx)
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"hot"}),
	}).Create(&videos).Error
}
