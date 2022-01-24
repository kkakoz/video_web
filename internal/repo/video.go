package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"video_web/internal/domain"
	"video_web/pkg/gormx"
	"video_web/pkg/model"
	"video_web/pkg/mysqlx"
)

var _ domain.IVideoRepo = (*VideoRepo)(nil)

type VideoRepo struct {
}

func NewVideoRepo() domain.IVideoRepo {
	return &VideoRepo{}
}

func (v *VideoRepo) AddVideo(ctx context.Context, video *domain.Video) error {
	db := mysqlx.GetDB(ctx)
	err := db.Create(video).Error
	return errors.Wrap(err, "添加失败")
}

func (v *VideoRepo) GetVideo(ctx context.Context, opts ...gormx.DBOption) (*domain.Video, error) {
	db := mysqlx.GetDB(ctx)
	video := &domain.Video{}
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Preload("Episodes").Find(video).Error
	return video, errors.Wrap(err, "查询失败")
}

func (v *VideoRepo) GetVideos(ctx context.Context, video *domain.Video, pager *model.Pager) ([]*domain.Video, error) {
	db := mysqlx.GetDB(ctx)
	db = pager.Paging(db)
	if video.Type != 0 {
		db = db.Where("type = ?", video.Type)
	}
	if video.Category != 0 {
		db = db.Where("category = ?", video.Category)
	}
	videos := make([]*domain.Video, 0)
	err := db.Find(&videos).Error
	return videos, err
}

func (v *VideoRepo) AddEpisode(ctx context.Context, episode *domain.Episode) error {
	db := mysqlx.GetDB(ctx)
	err := db.Create(episode).Error
	return errors.Wrap(err, "添加失败")
}

func (v *VideoRepo) GetEpisode(ctx context.Context, episodeId int64) (*domain.Episode, error) {
	db := mysqlx.GetDB(ctx)
	episode := &domain.Episode{}
	err := db.Where("id = ?", episodeId).Find(episode).Error
	return episode, errors.Wrap(err, "查询失败")
}

func (v *VideoRepo) GetEpisodes(ctx context.Context, videoId int64) ([]*domain.Episode, error) {
	db := mysqlx.GetDB(ctx)
	episodes := make([]*domain.Episode, 0)
	err := db.Where("video_id = ?", videoId).Order("order").Find(&episodes).Error
	return episodes, errors.Wrap(err, "查询失败")
}

func (v *VideoRepo) GetLastEpisode(ctx context.Context, videoId int64) (*domain.Episode, error) {
	db := mysqlx.GetDB(ctx)
	episode := &domain.Episode{}
	err := db.Where("video_id = ? and next_id = 0", videoId).Find(episode).Error
	return episode, errors.Wrap(err, "查询失败")
}

func (v *VideoRepo) UpdateEpisode(ctx context.Context, episode *domain.Episode) error {
	db := mysqlx.GetDB(ctx)
	err := db.Save(episode).Error
	return errors.Wrap(err, "更新失败")
}

func (v *VideoRepo) DelEpisode(ctx context.Context, episodeId int64) error {
	db := mysqlx.GetDB(ctx)
	err := db.Delete(&domain.Episode{ID: episodeId}).Error
	return errors.Wrap(err, "删除失败")
}

func (v *VideoRepo) UpdateAfterOrderEpisode(ctx context.Context, videoId int64, index int64, updateVal int) error {
	db := mysqlx.GetDB(ctx)
	err := db.Model(&domain.Episode{}).Where("video_id = ? and index >= ?", videoId, index).
		Update("index", gorm.Expr("index + ?", updateVal)).Error
	return errors.Wrap(err, "更新失败")
}