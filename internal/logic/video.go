package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opts"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"video_web/internal/domain"
	"video_web/pkg/local"
)

var _ domain.IVideoLogic = (*VideoLogic)(nil)

type VideoLogic struct {
	videoRepo   domain.IVideoRepo
	episodeRepo domain.IEpisodeRepo
}

func NewVideoLogic(videoRepo domain.IVideoRepo, episodeRepo domain.IEpisodeRepo) domain.IVideoLogic {
	return &VideoLogic{videoRepo: videoRepo, episodeRepo: episodeRepo}
}

func (v VideoLogic) AddVideo(ctx context.Context, video *domain.Video) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	video.UserId = user.ID
	return v.videoRepo.Add(ctx, video)
}

func (v VideoLogic) AddEpisode(ctx context.Context, episode *domain.Episode) (err error) {
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	err = v.videoRepo.UpdateAfterOrderEpisode(ctx, episode.VideoId, episode.Order, 1)
	if err != nil {
		return err
	}
	err = v.videoRepo.Updates(ctx, map[string]any{"episode_count": gorm.Expr("episode_count + 1")}, opts.Where("id = ?", episode.VideoId))
	if err != nil {
		return err
	}

	err = v.episodeRepo.Add(ctx, episode)
	return err
}

func (v VideoLogic) DelEpisode(ctx context.Context, episodeId int64) (err error) {
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	episode, err := v.episodeRepo.GetById(ctx, episodeId)
	if err != nil {
		return err
	}
	err = v.videoRepo.UpdateAfterOrderEpisode(ctx, episode.VideoId, episode.Order, -1)
	if err != nil {
		return err
	}
	err = v.videoRepo.Updates(ctx, map[string]any{"episode_count": gorm.Expr("episode_count - 1")}, opts.Where("id = ?", episode.VideoId))
	if err != nil {
		return err
	}
	return v.episodeRepo.DeleteById(ctx, episodeId)
}

// 获取视频详情
func (v VideoLogic) GetVideo(ctx context.Context, videoId int64) (*domain.Video, error) {
	video, err := v.videoRepo.GetById(ctx, videoId)
	if err != nil {
		return nil, err
	}
	video.Episodes, err = v.episodeRepo.GetList(ctx, opts.EQ("video_id", videoId))
	if err != nil {
		return nil, err
	}
	value := map[string]any{
		"view": gorm.Expr("view + 1"),
	}
	err = v.videoRepo.Updates(ctx, value, opts.Where("id = ?", videoId))
	if err != nil {
		return nil, err
	}
	return video, nil
}

// 获取视频列表
func (v VideoLogic) GetVideos(ctx context.Context, categoryId uint, lastValue uint, orderType uint8) ([]*domain.Video, error) {
	options := opts.NewOpts().Limit(15)
	if categoryId > 0 {
		options = options.Where("category_id = ?", categoryId)
	}
	options = lo.Ternary(orderType == 1,
		options.IsWhere(lastValue != 0, "id < ?", lastValue).Order("id desc"),     // 时间排序
		options.IsWhere(lastValue != 0, "view < ?", lastValue).Order("view desc")) // 热度排序
	return v.videoRepo.GetList(ctx, options...)
}
