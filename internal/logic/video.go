package logic

import (
	"context"
	"fmt"
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
		var duration int64
		resources := lo.Map(req.Resources, func(resource dto.Resource, i int) *entity.Resource {
			duration += resource.Duration
			return &entity.Resource{
				Name:     lo.Ternary(resource.Name != "", resource.Name, fmt.Sprintf("第%d集", i+1)),
				UserId:   user.ID,
				Duration: resource.Duration,
				Url:      resource.Url,
			}
		})

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
			Collect:    0,
			Duration:   duration,
			Hot:        0,
			Resources:  resources,
			State:      entity.VideoStateDefault,
			PublishAt:  req.PublishAt,
		}

		err = repo.Video().Add(ctx, video)
		if err != nil {
			return err
		}

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

	options = options.Limit(req.Pager.GetLimit()).Offset(req.Pager.GetOffset()).Preload("Resources").Preload("Category")

	if req.OrderType == 0 {
		options = options.Order("id desc") // 时间排行
	} else {
		options = options.Order("view desc") // 热度排序
	}
	list, err := repo.Video().GetList(ctx, options...)
	return list, count, err
}

func (item *videoLogic) Get(ctx context.Context, req *dto.VideoId) (*entity.Video, error) {
	return repo.Video().Get(ctx, opt.NewOpts().Preload("User").Preload("Resources").EQ("id", req.VideoId)...)
}

func (item *videoLogic) List(ctx context.Context, req *dto.Videos) ([]*entity.Video, error) {
	options := opt.NewOpts().Limit(10).Where("state = ?", entity.VideoStateNormal).
		IsWhere(req.CategoryId > 0, "category_id = ?", req.CategoryId).IsWhere(req.UserId != 0, "user_id = ?", req.UserId).
		Preload("User").Preload("Category")

	video, err := repo.Video().GetById(ctx, req.LastId)
	if err != nil {
		return nil, err
	}

	if req.LastId != 0 {
		options = options.Where("created_at <= ? and id < ?", video.CreatedAt, req.LastId)
	}
	options = options.Order("created_at desc, id desc") // 时间排行
	//if req.OrderType == 0 {
	//	if req.LastId != 0 {
	//		options = options.Where("created_at <= ? and id < ?", video.CreatedAt, req.LastId)
	//	}
	//	options = options.Order("created_at desc, id desc") // 时间排行
	//} else {
	//	if req.LastId != 0 {
	//		options = options.Where("view <= ? and id < ?", video.View, req.LastId)
	//	}
	//	options = options.Order("view desc, id desc") // 热度排序
	//}
	return repo.Video().GetList(ctx, options...)
}

func (item *videoLogic) UserState(ctx context.Context, req *dto.VideoId) (*vo.UserState, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
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

	collect, err := repo.Collect().GetExist(ctx, opt.Where("user_id = ? and target_id = ?", user.ID, video.ID))
	if err != nil {
		return nil, err
	}
	userState.UserCollect = collect

	followd, err := repo.Follow().GetExist(ctx, opt.Where("followed_user_id = ? and user_id = ?", video.UserId, user.ID))
	if err != nil {
		return nil, err
	}
	userState.FollowedCreator = followd

	haveNewsFeed, err := repo.Newsfeed().GetExist(ctx, opt.Where("user_id = ? and target_id = ? and target_type = ?", user.ID, video.ID, entity.NewsfeedTargetTypeVideo))
	if err != nil {
		return nil, err
	}
	userState.UserShared = haveNewsFeed

	return userState, nil
}

func (item *videoLogic) Recommend(ctx context.Context, req *dto.VideoId) ([]*entity.Video, error) {
	video, err := repo.Video().GetById(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	return repo.Video().GetList(ctx, opt.NewOpts().Where("category_id = ? and id != ?", video.CategoryId, video.ID).Preload("User").Order("hot desc").Limit(10)...)
}

func (item *videoLogic) Rankings(ctx context.Context, req *dto.Rankings) ([]*entity.Video, error) {
	options := opt.NewOpts().Limit(10).Where("state = ?", entity.VideoStateNormal).
		Preload("User").Preload("Category")

	video, err := repo.Video().GetById(ctx, req.LastId)
	if err != nil {
		return nil, err
	}

	if req.LastId != 0 {
		options = options.Where("view <= ? and id < ?", video.View, req.LastId)
	}
	options = options.Order("view desc, id desc") // 热度排序
	return repo.Video().GetList(ctx, options...)
}
