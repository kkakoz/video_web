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

func (videoLogic) AddVideo(ctx context.Context, req *dto.VideoAdd) (err error) {
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
		Duration:   req.Duration,
		State:      lo.Ternary(req.Show, entity.VideoStateNormal, entity.VideoStateHidden),
	}
	err = repo.Video().Add(ctx, episode)
	return err
}

func (videoLogic) DelVideo(ctx context.Context, req *dto.VideoId) (err error) {
	return repo.Video().DeleteById(ctx, req.VideoId)
}

// 获取视频详情
func (videoLogic) GetVideo(ctx context.Context, videoId int64) (*entity.Video, error) {
	video, err := repo.Video().Get(ctx, opt.NewOpts().Where("id = ? ", videoId).Preload("User")...)
	if err != nil {
		return nil, err
	}
	return video, nil
}

// 获取视频列表
func (videoLogic) GetVideos(ctx context.Context, categoryId uint, lastValue uint, orderType uint8) ([]*entity.Video, error) {
	options := opt.NewOpts().Limit(10).Where("state = ?", entity.VideoStateNormal).Preload("User")
	if categoryId > 0 {
		options = options.Where("category_id = ?", categoryId)
	}
	options = lo.Ternary(orderType == 0,
		options.IsWhere(lastValue != 0, "id < ?", lastValue).Order("id desc"),     // 时间排序
		options.IsWhere(lastValue != 0, "view < ?", lastValue).Order("view desc")) // 热度排序
	return repo.Video().GetList(ctx, options...)
}

func (videoLogic) GetPageList(ctx context.Context, req *dto.BackVideoList) ([]*entity.Video, int64, error) {
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

func (videoLogic) GetPageCollections(ctx context.Context, req *dto.BackCollectionList) ([]*entity.Collection, int64, error) {
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

func (videoLogic) GetCollection(ctx context.Context, req *dto.CollectionId) (*entity.Collection, error) {
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

func (videoLogic) UserState(ctx context.Context, req *dto.VideoId) (*vo.UserState, error) {
	user, ok := local.GetUserLocal(ctx)
	if !ok {
		return &vo.UserState{}, nil
	}

	video, err := Video().GetVideo(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	userState := &vo.UserState{}

	var like *entity.Like
	if video.CollectionId == 0 {
		like, err = repo.Like().Get(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, entity.LikeTargetTypeVideo, video.ID))
	} else {
		like, err = repo.Like().Get(ctx, opt.Where("user_id = ? and target_type = ? and target_id = ?", user.ID, entity.LikeTargetTypeCollection, video.CollectionId))
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
