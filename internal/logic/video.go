package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/local"
	"video_web/pkg/errno"
)

type videoLogic struct {
}

var videoOnce sync.Once
var _video *videoLogic

func Video() *videoLogic {
	videoOnce.Do(func() {
		_video = &videoLogic{}
	})
	return _video
}

func (item *videoLogic) AddCollection(ctx context.Context, req *dto.CollectionAdd) error {
	if req.Videos == nil || len(req.Videos) == 0 {
		return errno.New400("请选择视频草稿或者上传视频")
	}

	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		collection := &entity.Collection{}
		collection.UserName = user.Name
		err = copier.Copy(collection, req)
		if err != nil {
			return errors.WithStack(err)
		}
		collection.UserId = user.ID
		collection.EpisodeCount = int64(len(req.Videos))

		videos := lo.Map(req.Videos, func(video dto.VideoEasy, i int) *entity.Video {
			return &entity.Video{
				Name:         lo.Ternary(video.Name != "", video.Name, fmt.Sprintf("第%d集", i+1)),
				CategoryId:   req.CategoryId,
				Cover:        req.Cover,
				Brief:        req.Brief,
				UserId:       user.ID,
				UserName:     user.Name,
				UserAvatar:   user.Avatar,
				Url:          video.Url,
				CollectionId: collection.ID,
			}
		})
		collection.Videos = videos

		err = repo.Collection().Add(ctx, collection)
		if err != nil {
			return err
		}

		ids := lo.Map(videos, func(episode *entity.Video, i int) int64 {
			return episode.ID
		})

		collection.Orders = string(lo.Must(json.Marshal(ids)))
		err = repo.Collection().Updates(ctx, collection, opt.Where("id = ?", collection.ID))
		return err
	})
}

func (item *videoLogic) AddVideo(ctx context.Context, req *dto.VideoAdd) (err error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	episode := &entity.Video{
		Name:       req.Name,
		CategoryId: req.CategoryId,
		Cover:      req.Cover,
		Brief:      req.Brief,
		UserId:     user.ID,
		UserName:   user.Name,
		UserAvatar: user.Avatar,
		Url:        req.Url,
		Show:       req.Show,
		Duration:   req.Duration,
	}
	err = repo.Video().Add(ctx, episode)
	return err
}

func (item *videoLogic) DelVideo(ctx context.Context, req *dto.VideoId) (err error) {
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		return repo.Video().DeleteById(ctx, req.VideoId)
	})
}

// 获取视频详情
func (item *videoLogic) GetVideo(ctx context.Context, videoId int64) (*entity.Collection, error) {
	collection, err := repo.Collection().GetById(ctx, videoId)
	if err != nil {
		return nil, err
	}

	collection.Videos, err = repo.Video().GetList(ctx, opt.Where("collection = ?", collection.ID))
	if err != nil {
		return nil, err
	}
	value := map[string]any{
		"view": gorm.Expr("view + 1"),
	}
	err = repo.Video().Updates(ctx, value, opt.Where("id = ?", videoId))
	if err != nil {
		return nil, err
	}
	return collection, nil
}

// 获取视频列表
func (item *videoLogic) GetVideos(ctx context.Context, categoryId uint, lastValue uint, orderType uint8) ([]*entity.Video, error) {
	options := opt.NewOpts().Limit(15).Where("show = ?", true)
	if categoryId > 0 {
		options = options.Where("category_id = ?", categoryId)
	}
	options = lo.Ternary(orderType == 0,
		options.IsWhere(lastValue != 0, "id < ?", lastValue).Order("id desc"),     // 时间排序
		options.IsWhere(lastValue != 0, "view < ?", lastValue).Order("view desc")) // 热度排序
	return repo.Video().GetList(ctx, options...)
}

func (item *videoLogic) GetPageList(ctx context.Context, req *dto.BackVideoList) ([]*entity.Video, int64, error) {
	options := opt.NewOpts().IsLike(req.Name != "", "name", req.Name)
	if req.CategoryId > 0 {
		options = options.Where("category_id = ?", req.CategoryId)
	}
	count, err := repo.Video().Count(ctx, options...)
	if err != nil {
		return nil, 0, err
	}

	options = options.Limit(req.Pager.GetLimit()).Offset(req.Pager.GetOffset())

	options = lo.Ternary(req.OrderType == 0,
		options.Order("id desc"),   // 时间排序
		options.Order("view desc")) // 热度排序
	list, err := repo.Video().GetList(ctx, options...)
	return list, count, err
}

func (item *videoLogic) GetPageCollections(ctx context.Context, req *dto.BackCollectionList) ([]*entity.Collection, int64, error) {
	options := opt.NewOpts().IsLike(req.Name != "", "name", req.Name)
	count, err := repo.Collection().Count(ctx, options...)
	if err != nil {
		return nil, 0, err
	}

	options = options.Limit(req.Pager.GetLimit()).Offset(req.Pager.GetOffset()).Preload("Videos")

	options = lo.Ternary(req.OrderType == 0,
		options.Order("id desc"),   // 时间排序
		options.Order("view desc")) // 热度排序
	list, err := repo.Collection().GetList(ctx, options...)
	return list, count, err
}

func (item *videoLogic) GetCollection(ctx context.Context, req *dto.CollectionId) (*entity.Collection, error) {
	return repo.Collection().Get(ctx, opt.NewOpts().EQ("id", req.CollectionId).Preload("Videos")...)
}
