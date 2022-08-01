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

type collectionLogic struct {
}

var collectionOnce sync.Once
var _collection *collectionLogic

func Collection() *collectionLogic {
	collectionOnce.Do(func() {
		_collection = &collectionLogic{}
	})
	return _collection
}

func (item *collectionLogic) Add(ctx context.Context, req *dto.CollectionAdd) error {
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

		videos := lo.Map(req.Videos, func(video dto.VideoEasy, i int) *entity.Video {
			duration += video.Duration
			return &entity.Video{
				Name:       lo.Ternary(video.Name != "", video.Name, fmt.Sprintf("第%d集", i+1)),
				CategoryId: req.CategoryId,
				Cover:      req.Cover,
				Brief:      req.Brief,
				UserId:     user.ID,
				UserName:   user.Name,
				Duration:   video.Duration,
				UserAvatar: user.Avatar,
				Url:        video.Url,
				State:      lo.Ternary(i == 0, entity.VideoStateNormal, entity.VideoStateHidden),
			}
		})

		collection := &entity.Collection{
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
			UserId:     user.ID,
			UserName:   user.Name,
			UserAvatar: user.Avatar,
			State:      req.State,
			PublishAt:  req.PublishAt,
			//CreatedAt:  timex.Time{},
			//UpdatedAt:  timex.Time{},
			Orders: "",
			Videos: videos,
		}

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

func (item *collectionLogic) GetPageList(ctx context.Context, req *dto.BackCollectionList) ([]*entity.Collection, int64, error) {
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

func (item *collectionLogic) Get(ctx context.Context, req *dto.CollectionId) (*entity.Collection, error) {
	collection, err := repo.Collection().Get(ctx, opt.NewOpts().EQ("id", req.CollectionId)...)
	if err != nil {
		return nil, err
	}
	list, err := repo.Video().GetList(ctx, opt.Where("collection_id = ?", collection.ID))
	idVideoMap := lo.GroupBy(list, func(v *entity.Video) int64 {
		return v.ID
	})
	ids := make([]int64, 0)
	err = json.Unmarshal([]byte(collection.Orders), &ids)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	orderList := make([]*entity.Video, 0)
	for _, videos := range idVideoMap {
		for _, video := range videos {
			orderList = append(orderList, video)
		}
	}

	collection.Videos = orderList
	return collection, nil
}
