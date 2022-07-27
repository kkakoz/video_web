package repo

import (
	"context"
	"sync"
	"video_web/internal/model"

	"github.com/kkakoz/ormx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type collectionRepo struct {
	ormx.IRepo[model.Collection]
}

var collectionOnce sync.Once
var _collection *collectionRepo

func Collection() *collectionRepo {
	collectionOnce.Do(func() {
		_collection = &collectionRepo{IRepo: ormx.NewRepo[model.Collection]()}
	})
	return _collection
}

func (v *collectionRepo) UpdateAfterOrderEpisode(ctx context.Context, videoId int64, index int64, updateVal int) error {
	db := ormx.DB(ctx)
	err := db.Model(&model.Video{}).Where("video_id = ? and index >= ?", videoId, index).
		Update("index", gorm.Expr("index + ?", updateVal)).Error
	return errors.Wrap(err, "更新失败")
}
