package logic

import (
	"context"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/pkg/model"
	"github/kkakoz/video_web/pkg/mysqlx"
)

type VideoLogic struct {
	videoRepo  domain.IVideoRepo
}

func (v VideoLogic) AddVideo(ctx context.Context, video *domain.Video) error {
	return v.videoRepo.AddVideo(ctx, video)
}

func (v VideoLogic) AddEpisode(ctx context.Context, episode *domain.Episode) (err error) {
	ctx, checkErr := mysqlx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	lastEpisode, err := v.videoRepo.GetLastEpisode(ctx, episode.VideoId)
	if err != nil {
		return err
	}
	// 添加最后一个
	episode.PreId = lastEpisode.ID
	err = v.videoRepo.AddEpisode(ctx, episode)
	if err != nil {
		return err
	}
	lastEpisode.NextId = episode.ID
	return v.videoRepo.UpdateEpisode(ctx, episode)
}

func (v VideoLogic) DelEpisode(ctx context.Context, episodeId int64) (err error) {
	ctx, checkErr := mysqlx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	episode, err := v.videoRepo.GetEpisode(ctx, episodeId)
	if err != nil {
		return err
	}
	if episode.PreId != 0 {
		preEpisode, err := v.videoRepo.GetEpisode(ctx, episode.PreId)
		if err != nil {
			return err
		}
		preEpisode.NextId = episode.NextId
		err = v.videoRepo.UpdateEpisode(ctx, preEpisode)
		if err != nil {
			return err
		}
	}
	if episode.NextId != 0 {
		nextEpisode, err := v.videoRepo.GetEpisode(ctx, episode.NextId)
		if err != nil {
			return err
		}
		nextEpisode.PreId = episode.PreId
		err = v.videoRepo.UpdateEpisode(ctx, nextEpisode)
		if err != nil {
			return err
		}
	}
	return v.videoRepo.DelEpisode(ctx, episodeId)
}

func (v VideoLogic) GetVideo(ctx context.Context, videoId int64) (*domain.Video, error) {
	video, err := v.videoRepo.GetVideo(ctx, videoId)
	if err != nil {
		return video, err
	}
	episodes, err := v.videoRepo.GetEpisodes(ctx, videoId)
	if err != nil {
		return video, err
	}
	// todo episodes排序
	video.Episodes = episodes
	return video, nil
}

func (v VideoLogic) GetVideos(ctx context.Context, video *domain.Video, pager *model.Pager) ([]*domain.Video, error) {
	return v.videoRepo.GetVideos(ctx, video, pager)
}
