package repo

import (
	"context"
	"github.com/samber/lo"
	"video_web/internal/model"

	"github.com/kkakoz/ormx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type VideoRepo struct {
	ormx.IRepo[model.Video]
}

func NewVideoRepo() *VideoRepo {
	return &VideoRepo{
		ormx.NewRepo[model.Video](),
	}
}

func (v *VideoRepo) UpdateAfterOrderEpisode(ctx context.Context, videoId int64, index int64, updateVal int) error {
	db := ormx.DB(ctx)
	err := db.Model(&model.Episode{}).Where("video_id = ? and index >= ?", videoId, index).
		Update("index", gorm.Expr("index + ?", updateVal)).Error
	return errors.Wrap(err, "更新失败")
}

func (v *VideoRepo) AddEpisode(ctx context.Context, videoId int64, episodeIds []int64) error {
	db := ormx.DB(ctx)
	list := lo.Map(episodeIds, func(episodeId int64, i int) *model.VideoEpisode {
		return &model.VideoEpisode{
			VideoId:   videoId,
			EpisodeId: episodeId,
		}
	})
	err := db.Create(&list).Error
	return errors.WithStack(err)
}

func (v *VideoRepo) GetEpisodeIds(ctx context.Context, videoId int64) ([]int64, error) {
	db := ormx.DB(ctx)
	list := []int64{}
	err := db.Model(&model.VideoEpisode{}).Where("video_id = ?", videoId).Pluck("episode_id", &list).Error
	return list, errors.WithStack(err)
}

func (v *VideoRepo) DeleteEpisode(ctx context.Context, videoId int64, episodeId int64) error {
	db := ormx.DB(ctx)
	return db.Where("video_id = ? and episode_id = ?", videoId, episodeId).Delete(&model.VideoEpisode{}).Error
}
