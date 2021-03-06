package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"video_web/internal/consts"
	"video_web/internal/dto/request"
	"video_web/internal/model"
	"video_web/internal/repo"
	"video_web/pkg/errno"
	"video_web/pkg/local"
)

type VideoLogic struct {
	videoRepo   *repo.VideoRepo
	episodeRepo *repo.EpisodeRepo
}

func NewVideoLogic(videoRepo *repo.VideoRepo, episodeRepo *repo.EpisodeRepo) *VideoLogic {
	return &VideoLogic{videoRepo: videoRepo, episodeRepo: episodeRepo}
}

func (item *VideoLogic) Add(ctx context.Context, req *request.VideoAddReq) error {
	if req.Episodes == nil || len(req.Episodes) == 0 {
		return errno.New400("请选择视频草稿或者上传视频")
	}

	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	video := &model.Video{}
	video.UserName = user.Name
	err = copier.Copy(video, req)
	if err != nil {
		return errors.WithStack(err)
	}
	video.UserId = user.ID

	video.EpisodeCount = int64(len(req.EpisodeIds))
	err = item.videoRepo.Add(ctx, video)
	if err != nil {
		return err
	}

	list := lo.Map(req.Episodes, func(episode request.EpisodeEasy, i int) *model.Episode {
		return &model.Episode{
			Name:   lo.Ternary(episode.Name != "", episode.Name, fmt.Sprintf("第%d集", i)),
			UserId: user.ID,
			Url:    episode.Url,
		}
	})
	err = item.episodeRepo.AddList(ctx, list)
	if err != nil {
		return err
	}
	ids := lo.Map(list, func(episode *model.Episode, i int) int64 {
		return episode.ID
	})
	err = item.videoRepo.AddEpisode(ctx, video.ID, ids)
	return err

}

func (item *VideoLogic) AddEpisode(ctx context.Context, req *request.EpisodeAddReq) (err error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	episode := &model.Episode{
		Name:   req.Name,
		UserId: user.ID,
		Url:    req.Url,
	}
	err = item.episodeRepo.Add(ctx, episode)
	if err != nil {
		return err
	}
	if req.VideoId != 0 { // 给合集添加新的一集
		return item.videoRepo.AddEpisode(ctx, req.VideoId, []int64{episode.ID})
	}

	if req.AddType == consts.VideoTypeSingle { // 单独添加为视频
		return item.Add(ctx, &request.VideoAddReq{
			Name:       req.Name,
			Type:       req.AddType,
			CategoryId: req.CategoryId,
			Cover:      req.Cover,
			Brief:      req.Brief,
			EpisodeIds: []int64{episode.ID},
		})
	}
	return nil
}

func (item *VideoLogic) DelVideoEpisode(ctx context.Context, req *request.EpisodeIdReq) (err error) {
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	err = item.videoRepo.DeleteEpisode(ctx, req.VideoId, req.EpisodeId)
	if err != nil {
		return err
	}
	err = item.videoRepo.Updates(ctx, map[string]any{"episode_count": gorm.Expr("episode_count - 1")}, opt.Where("id = ?", req.VideoId))
	if err != nil {
		return err
	}
	return item.episodeRepo.DeleteById(ctx, req.EpisodeId)
}

// 获取视频详情
func (item *VideoLogic) GetVideo(ctx context.Context, videoId int64) (*model.Video, error) {
	video, err := item.videoRepo.GetById(ctx, videoId)
	if err != nil {
		return nil, err
	}
	ids, err := item.videoRepo.GetEpisodeIds(ctx, videoId)
	if err != nil {
		return nil, err
	}
	video.Episodes, err = item.episodeRepo.GetList(ctx, opt.In("id", ids))
	if err != nil {
		return nil, err
	}
	value := map[string]any{
		"view": gorm.Expr("view + 1"),
	}
	err = item.videoRepo.Updates(ctx, value, opt.Where("id = ?", videoId))
	if err != nil {
		return nil, err
	}
	return video, nil
}

// 获取视频列表
func (item *VideoLogic) GetVideos(ctx context.Context, categoryId uint, lastValue uint, orderType uint8) ([]*model.Video, error) {
	options := opt.NewOpts().Limit(15)
	if categoryId > 0 {
		options = options.Where("category_id = ?", categoryId)
	}
	options = lo.Ternary(orderType == 0,
		options.IsWhere(lastValue != 0, "id < ?", lastValue).Order("id desc"),     // 时间排序
		options.IsWhere(lastValue != 0, "view < ?", lastValue).Order("view desc")) // 热度排序
	return item.videoRepo.GetList(ctx, options...)
}

func (item *VideoLogic) GetBackVideos(ctx context.Context, categoryId uint, orderType uint8, pager request.Pager) ([]*model.Video, int64, error) {
	options := opt.NewOpts().Limit(15)
	if categoryId > 0 {
		options = options.Where("category_id = ?", categoryId)
	}
	count, err := item.videoRepo.Count(ctx, options...)
	if err != nil {
		return nil, 0, err
	}

	options = options.Limit(pager.GetLimit()).Offset(pager.GetOffset())

	options = lo.Ternary(orderType == 0,
		options.Order("id desc"),   // 时间排序
		options.Order("view desc")) // 热度排序
	list, err := item.videoRepo.GetList(ctx, options...)
	return list, count, err
}
