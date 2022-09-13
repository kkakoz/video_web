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
	"video_web/internal/model/vo"
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
				Name:     lo.Ternary(video.Name != "", video.Name, fmt.Sprintf("第%d集", i+1)),
				Cover:    req.Cover,
				Brief:    req.Brief,
				Duration: video.Duration,
				Url:      video.Url,
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

func (collectionLogic) GetPageList(ctx context.Context, req *dto.BackCollectionList) ([]*entity.Collection, int64, error) {
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

func (collectionLogic) Get(ctx context.Context, req *dto.CollectionId) (*entity.Collection, error) {
	collection, err := repo.Collection().GetById(ctx, req.CollectionId)
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

func (collectionLogic) UserState(ctx context.Context, req *dto.CollectionId) (*vo.UserState, error) {
	user, ok := local.GetUserLocal(ctx)
	if !ok {
		return &vo.UserState{}, nil
	}

	collection, err := Collection().Get(ctx, req)
	if err != nil {
		return nil, err
	}

	userState := &vo.UserState{}

	var like *entity.Like
	like, err = repo.Like().Get(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, entity.LikeTargetTypeCollection, collection.ID))
	if err != nil {
		return nil, err
	}
	if like != nil {
		if like.Like {
			userState.UserLike = true
		} else {
			userState.UserDisLike = true
		}
	}

	exist, err := repo.Follow().GetExist(ctx, opt.Where("followed_user_id = ? and user_id = ?", collection.UserId, user.ID))
	if exist {
		userState.FollowedCreator = true
	}

	return userState, nil
}

func (item *collectionLogic) List(ctx context.Context, req *dto.CollectionList) ([]*entity.Collection, error) {
	options := opt.NewOpts().Limit(10).Where("state = ?", entity.CollectionStateNormal).Preload("User")
	if req.CategoryId > 0 {
		options = options.Where("category_id = ?", req.CategoryId)
	}
	options = lo.Ternary(req.OrderType == 0,
		options.IsWhere(req.LastValue != 0, "id < ?", req.LastValue).Order("id desc"),     // 时间排序
		options.IsWhere(req.LastValue != 0, "view < ?", req.LastValue).Order("view desc")) // 热度排序
	return repo.Collection().GetList(ctx, options...)
}
