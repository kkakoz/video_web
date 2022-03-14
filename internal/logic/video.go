package logic

import (
	"context"
	"video_web/internal/domain"
	"video_web/pkg/local"
	"video_web/pkg/model"

	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opts"
)

var _ domain.IVideoLogic = (*VideoLogic)(nil)

type VideoLogic struct {
	videoRepo   domain.IVideoRepo
	episodeRepo domain.IEpisodeRepo
}

func NewVideoLogic(videoRepo domain.IVideoRepo) domain.IVideoLogic {
	return &VideoLogic{videoRepo: videoRepo}
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
	return v.episodeRepo.DeleteById(ctx, episodeId)
}

func (v VideoLogic) GetVideo(ctx context.Context, videoId int64) (*domain.Video, error) {
	video, err := v.videoRepo.GetById(ctx, videoId)
	if err != nil {
		return nil, err
	}
	episodes, err := v.episodeRepo.GetList(ctx, opts.EQ("video_id", videoId))
	if err != nil {
		return nil, err
	}
	video.Episodes = episodes
	return video, nil
}

func (v VideoLogic) GetVideos(ctx context.Context, video *domain.Video, pager *model.Pager) ([]*domain.Video, error) {
	return v.videoRepo.GetList(ctx, opts.NewOpts().Page(pager.Page, pager.PageSize)...)
}
