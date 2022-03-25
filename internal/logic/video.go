package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opts"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"video_web/internal/consts"
	"video_web/internal/domain"
	"video_web/internal/dto/request"
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

func (v VideoLogic) Add(ctx context.Context, req *request.AddVideoReq) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	video := &domain.Video{}
	err = copier.Copy(video, req)
	video.UserId = user.ID
	video.EpisodeCount = int64(len(req.EpisodeIds))
	err = v.videoRepo.Add(ctx, video)
	if err != nil {
		return err
	}
	err = v.videoRepo.AddEpisode(ctx, video.ID, req.EpisodeIds)
	return err
}

func (v VideoLogic) AddEpisode(ctx context.Context, req *request.AddEpisodeReq) (err error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	episode := &domain.Episode{
		Name:   req.Name,
		UserId: user.ID,
		Url:    req.Url,
	}
	err = v.episodeRepo.Add(ctx, episode)
	if err != nil {
		return err
	}

	if req.AddType == consts.VideoTypeSingle {
		return v.Add(ctx, &request.AddVideoReq{
			Name:       req.Name,
			Type:       req.AddType,
			Category:   req.CategoryId,
			Cover:      req.Cover,
			Brief:      req.Brief,
			EpisodeIds: []int64{episode.ID},
		})
	}

	return nil
}

func (v VideoLogic) DelVideoEpisode(ctx context.Context, req *request.EpisodeIdReq) (err error) {
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	err = v.videoRepo.DeleteEpisode(ctx, req.VideoId, req.EpisodeId)
	if err != nil {
		return err
	}
	err = v.videoRepo.Updates(ctx, map[string]any{"episode_count": gorm.Expr("episode_count - 1")}, opts.Where("id = ?", req.VideoId))
	if err != nil {
		return err
	}
	return v.episodeRepo.DeleteById(ctx, req.EpisodeId)
}

// 获取视频详情
func (v VideoLogic) GetVideo(ctx context.Context, videoId int64) (*domain.Video, error) {
	video, err := v.videoRepo.GetById(ctx, videoId)
	if err != nil {
		return nil, err
	}
	ids, err := v.videoRepo.GetEpisodeIds(ctx, videoId)
	if err != nil {
		return nil, err
	}
	video.Episodes, err = v.episodeRepo.GetList(ctx, opts.In("id", ids))
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
