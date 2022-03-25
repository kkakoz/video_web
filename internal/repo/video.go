package repo

import (
	"context"
	"github.com/samber/lo"
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

func (v *VideoRepo) AddEpisode(ctx context.Context, videoId int64, episodeIds []int64) error {
	db := ormx.DB(ctx)
	list := lo.Map(episodeIds, func(episodeId int64, i int) *domain.VideoEpisode {
		return &domain.VideoEpisode{
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
	err := db.Model(&domain.VideoEpisode{}).Where("video_id = ?", videoId).Pluck("episode_id", &list).Error
	return list, errors.WithStack(err)
}

func (v *VideoRepo) DeleteEpisode(ctx context.Context, videoId int64, episodeId int64) error {
	db := ormx.DB(ctx)
	return errors.WithStack(db.Where("video_id = ? and episode_id = ?", episodeId).Delete(&domain.VideoEpisode{}).Error)
}
