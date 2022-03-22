package repo

import (
	"context"
	"video_web/internal/domain"

	"github.com/kkakoz/ormx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var _ domain.IVideoRepo = (*VideoRepo)(nil)

type VideoRepo struct {
	ormx.IRepo[domain.Video]
}

func NewVideoRepo() domain.IVideoRepo {
	return &VideoRepo{
		ormx.NewRepo[domain.Video](),
	}
}

func (v *VideoRepo) UpdateAfterOrderEpisode(ctx context.Context, videoId int64, index int64, updateVal int) error {
	db := ormx.DB(ctx)
	err := db.Model(&domain.Episode{}).Where("video_id = ? and index >= ?", videoId, index).
		Update("index", gorm.Expr("index + ?", updateVal)).Error
	return errors.Wrap(err, "更新失败")
}
