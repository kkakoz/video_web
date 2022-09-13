package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/model/vo"
	"video_web/internal/pkg/local"
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

func (item *videoLogic) Add(ctx context.Context, req *dto.VideoAdd) (*entity.Video, error) {

	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	var video *entity.Video
	err = ormx.Transaction(ctx, func(ctx context.Context) error {
		//videos := lo.Map(req.Videos, func(video dto.VideoEasy, i int) *entity.Resource {
		//	duration += video.Duration
		//	return &entity.Resource{
		//		Name:       lo.Ternary(video.Name != "", video.Name, fmt.Sprintf("第%d集", i+1)),
		//		UserId:     user.ID,
		//		UserName:   user.Name,
		//		Duration:   video.Duration,
		//		UserAvatar: user.Avatar,
		//		Url:        video.Url,
		//	}
		//})

		video = &entity.Video{
			Name:       req.Name,
			Type:       req.Type,
			UserId:     user.ID,
			CategoryId: req.CategoryId,
			Cover:      req.Cover,
			Brief:      req.Brief,
			View:       0,
			Like:       0,
			Comment:    0,
			Favorite:   0,
			Duration:   0,
			Hot:        0,
			//VideoCount: int64(len(req.Videos)),
			State:     entity.VideoStateDefault,
			PublishAt: req.PublishAt,
		}

		err = repo.Video().Add(ctx, video)
		if err != nil {
			return err
		}

		//ids := lo.Map(videos, func(episode *entity.Resource, i int) int64 {
		//	return episode.ID
		//})

		//video.Orders = string(lo.Must(json.Marshal(ids)))

		//err = repo.Video().Updates(ctx, video, opt.Where("id = ?", video.ID))
		return err
	})
	return video, err
}

func (videoLogic) GetPageList(ctx context.Context, req *dto.BackCollectionList) ([]*entity.Video, int64, error) {
	options := opt.NewOpts().IsLike(req.Name != "", "name", req.Name)
	count, err := repo.Video().Count(ctx, options...)
	if err != nil {
		return nil, 0, err
	}

	options = options.Limit(req.Pager.GetLimit()).Offset(req.Pager.GetOffset()).Preload("Resources")

	options = lo.Ternary(req.OrderType == 0,
		options.Order("id desc"), // 时间排序
		options.Order("view desc")) // 热度排序
	list, err := repo.Video().GetList(ctx, options...)
	return list, count, err
}

func (videoLogic) Get(ctx context.Context, req *dto.VideoId) (*entity.Video, error) {
	collection, err := repo.Video().Get(ctx, opt.NewOpts().EQ("id", req.VideoId)...)
	if err != nil {
		return nil, err
	}
	list, err := repo.Resource().GetList(ctx, opt.Where("collection_id = ?", collection.ID), opt.Order("order"))

	collection.Resources = list
	return collection, nil
}

func (item *videoLogic) List(ctx context.Context, req *dto.Videos) ([]*entity.Video, error) {
	options := opt.NewOpts().Limit(10).Where("state = ?", entity.VideoStateNormal).
		IsWhere(req.CategoryId > 0, "category_id = ?", req.CategoryId).IsWhere(req.UserId != 0, "user_id = ?", req.UserId).
		Preload("User")
	
	options = lo.Ternary(req.OrderType == 0,
		options.IsWhere(req.LastValue != 0, "id < ?", req.LastValue != 0).Order("id desc"), // 时间排序
		options.IsWhere(req.LastValue != 0, "view < ?", req.LastValue != 0).Order("view desc")) // 热度排序
	return repo.Video().GetList(ctx, options...)
}

func (videoLogic) UserState(ctx context.Context, req *dto.VideoId) (*vo.UserState, error) {
	user, ok := local.GetUserLocal(ctx)
	if !ok {
		return &vo.UserState{}, nil
	}

	video, err := repo.Video().GetById(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	userState := &vo.UserState{}

	like, err := repo.Like().Get(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, entity.LikeTargetTypeVideo, video.ID))
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

	exist, err := repo.Follow().GetExist(ctx, opt.Where("followed_user_id = ? and user_id = ?", video.UserId, user.ID))
	if exist {
		userState.FollowedCreator = true
	}

	return userState, nil
}
