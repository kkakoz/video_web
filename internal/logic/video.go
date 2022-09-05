package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/pkg/errors"
	"github.com/samber/lo"
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

func (item *videoLogic) Add(ctx context.Context, req *dto.CollectionAdd) error {
	if req.Videos == nil || len(req.Videos) == 0 {
		return errno.New400("请选择视频草稿或者上传视频")
	}

	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	return ormx.Transaction(ctx, func(ctx context.Context) error {

		// 总时长
		var duration int64 = 0

		videos := lo.Map(req.Videos, func(video dto.VideoEasy, i int) *entity.Resource {
			duration += video.Duration
			return &entity.Resource{
				Name:       lo.Ternary(video.Name != "", video.Name, fmt.Sprintf("第%d集", i+1)),
				UserId:     user.ID,
				UserName:   user.Name,
				Duration:   video.Duration,
				UserAvatar: user.Avatar,
				Url:        video.Url,
			}
		})

		collection := &entity.Video{
			Name:       req.Name,
			Type:       1,
			CategoryId: req.CategoryId,
			Cover:      req.Cover,
			Brief:      req.Brief,
			View:       0,
			Like:       0,
			Comment:    0,
			Favorite:   0,
			Duration:   duration,
			Hot:        0,
			VideoCount: int64(len(req.Videos)),
			State:      req.State,
			PublishAt:  req.PublishAt,
			//CreatedAt:  timex.Time{},
			//UpdatedAt:  timex.Time{},
			Orders:    "",
			Resources: videos,
		}

		err = repo.Video().Add(ctx, collection)
		if err != nil {
			return err
		}

		ids := lo.Map(videos, func(episode *entity.Resource, i int) int64 {
			return episode.ID
		})

		collection.Orders = string(lo.Must(json.Marshal(ids)))

		err = repo.Video().Updates(ctx, collection, opt.Where("id = ?", collection.ID))
		return err
	})
}

func (videoLogic) GetPageList(ctx context.Context, req *dto.BackCollectionList) ([]*entity.Video, int64, error) {
	options := opt.NewOpts().IsLike(req.Name != "", "name", req.Name)
	count, err := repo.Video().Count(ctx, options...)
	if err != nil {
		return nil, 0, err
	}

	options = options.Limit(req.Pager.GetLimit()).Offset(req.Pager.GetOffset()).Preload("Resources")

	options = lo.Ternary(req.OrderType == 0,
		options.Order("id desc"),   // 时间排序
		options.Order("view desc")) // 热度排序
	list, err := repo.Video().GetList(ctx, options...)
	return list, count, err
}

func (videoLogic) Get(ctx context.Context, req *dto.VideoId) (*entity.Video, error) {
	collection, err := repo.Video().Get(ctx, opt.NewOpts().EQ("id", req.VideoId)...)
	if err != nil {
		return nil, err
	}
	list, err := repo.Resource().GetList(ctx, opt.Where("collection_id = ?", collection.ID))
	idVideoMap := lo.GroupBy(list, func(v *entity.Resource) int64 {
		return v.ID
	})
	ids := make([]int64, 0)
	err = json.Unmarshal([]byte(collection.Orders), &ids)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	orderList := make([]*entity.Resource, 0)
	for _, videos := range idVideoMap {
		for _, video := range videos {
			orderList = append(orderList, video)
		}
	}

	collection.Resources = orderList
	return collection, nil
}

func (item *videoLogic) List(ctx context.Context, req *dto.Videos) ([]*entity.Video, error) {
	options := opt.NewOpts().Limit(10).Where("state = ?", entity.VideoStateNormal).Preload("User")

	options = options.IsWhere(req.CategoryId > 0, "category_id = ?", req.CategoryId)
	options = lo.Ternary(req.OrderType == 0,
		options.IsWhere(req.LastValue != 0, "id < ?", req.LastValue != 0).Order("id desc"),     // 时间排序
		options.IsWhere(req.LastValue != 0, "view < ?", req.LastValue != 0).Order("view desc")) // 热度排序
	return repo.Video().GetList(ctx, options...)
}
