package logic

import (
	"context"
	"encoding/json"
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
)

type resourceLogic struct {
}

var resourceOnce sync.Once
var _resource *resourceLogic

func Resource() *resourceLogic {
	resourceOnce.Do(func() {
		_resource = &resourceLogic{}
	})
	return _resource
}

func (resourceLogic) AddVideo(ctx context.Context, req *dto.ResourceAdd) (err error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	ctx, checkErr := ormx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	episode := &entity.Resource{
		Name:       req.Name,
		Brief:      req.Brief,
		UserId:     user.ID,
		UserName:   user.Name,
		UserAvatar: user.Avatar,
		Url:        req.Url,
		Duration:   req.Duration,
	}
	err = repo.Resource().Add(ctx, episode)
	return err
}

func (resourceLogic) DelVideo(ctx context.Context, req *dto.ResourceId) (err error) {
	return repo.Resource().DeleteById(ctx, req.ResourceId)
}

// 获取视频详情
func (resourceLogic) Get(ctx context.Context, resourceId int64) (*entity.Resource, error) {
	video, err := repo.Resource().Get(ctx, opt.NewOpts().Where("id = ? ", resourceId).Preload("User")...)
	if err != nil {
		return nil, err
	}
	return video, nil
}

// 获取视频列表
func (resourceLogic) GetVideos(ctx context.Context, categoryId uint, lastValue uint, orderType uint8) ([]*entity.Resource, error) {
	options := opt.NewOpts().Limit(10).Where("state = ?", entity.VideoStateNormal).Preload("User")
	if categoryId > 0 {
		options = options.Where("category_id = ?", categoryId)
	}
	options = lo.Ternary(orderType == 0,
		options.IsWhere(lastValue != 0, "id < ?", lastValue).Order("id desc"),     // 时间排序
		options.IsWhere(lastValue != 0, "view < ?", lastValue).Order("view desc")) // 热度排序
	return repo.Resource().GetList(ctx, options...)
}

func (resourceLogic) GetPageList(ctx context.Context, req *dto.BackResourceList) ([]*entity.Resource, int64, error) {
	options := opt.NewOpts().IsLike(req.Name != "", "name", req.Name)
	options = options.IsWhere(req.CategoryId > 0, "category_id = ?", req.CategoryId)
	count, err := repo.Resource().Count(ctx, options...)
	if err != nil {
		return nil, 0, err
	}

	options = options.Limit(req.Pager.GetLimit()).Offset(req.Pager.GetOffset())

	options = lo.Ternary(req.OrderType == 0,
		options.Order("id desc"),   // 时间排序
		options.Order("view desc")) // 热度排序
	list, err := repo.Resource().GetList(ctx, options...)
	return list, count, err
}

func (resourceLogic) GetList(ctx context.Context, req *dto.VideoId) (*entity.Video, error) {
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

func (resourceLogic) UserState(ctx context.Context, req *dto.ResourceId) (*vo.UserState, error) {
	user, ok := local.GetUserLocal(ctx)
	if !ok {
		return &vo.UserState{}, nil
	}

	video, err := Resource().Get(ctx, req.ResourceId)
	if err != nil {
		return nil, err
	}

	userState := &vo.UserState{}

	var like *entity.Like
	if video.VideoId == 0 {
		like, err = repo.Like().Get(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, entity.LikeTargetTypeVideo, video.ID))
	} else {
		like, err = repo.Like().Get(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, entity.LikeTargetTypeCollection, video.VideoId))
	}
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
