package logic

import (
	"context"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/model/vo"
	"video_web/internal/pkg/local"
)

type newsfeed struct {
}

func Newsfeed() *newsfeed {
	return &newsfeed{}
}

func (newsfeed) Add(ctx context.Context, req *dto.NewsfeedAdd) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	err = repo.Newsfeed().Add(ctx, &entity.Newsfeed{
		UserId:     user.ID,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
		Content:    req.Content,
	})
	return err
}

func (us newsfeed) UserNews(ctx context.Context, req *dto.UserId) ([]*vo.NewsFeedVo, error) {
	list, err := repo.Newsfeed().GetList(ctx, opt.NewOpts().Where("user_id = ?", req.UserId).Order("created_at desc")...)
	if err != nil {
		return nil, err
	}
	return us.dealNewsFeed(ctx, list)
}

func (us newsfeed) List(ctx context.Context, req *dto.NewsfeedList) ([]*vo.NewsFeedVo, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	followerIds := make([]int64, 0)
	err = repo.Follow().Pluck(ctx, "id", followerIds, opt.NewOpts().Where("user_id = ?", user.ID)...)
	if err != nil {
		return nil, err
	}

	list, err := repo.Newsfeed().GetList(ctx, opt.NewOpts().Where("user_id in ?", followerIds).Order("created_at desc")...)

	return us.dealNewsFeed(ctx, list)
}

func (newsfeed) dealNewsFeed(ctx context.Context, list []*entity.Newsfeed) ([]*vo.NewsFeedVo, error) {
	videoIds := make([]int64, 0)
	for _, news := range list {
		if news.TargetType == entity.NewsfeedTargetTypeVideo {
			videoIds = append(videoIds, news.TargetId)
		}
	}
	videos, err := repo.Video().GetList(ctx, opt.NewOpts().In("id", videoIds).Preload("Resources")...)
	if err != nil {
		return nil, err
	}
	videoMap := lo.GroupBy(videos, func(v *entity.Video) int64 {
		return v.ID
	})

	newsFeeds := lo.Map(list, func(v *entity.Newsfeed, i int) *vo.NewsFeedVo {
		target, ok := videoMap[v.TargetId]
		return &vo.NewsFeedVo{
			ID:         v.ID,
			UserId:     v.UserId,
			TargetType: v.TargetType,
			TargetId:   v.TargetId,
			Content:    v.Content,
			CreatedAt:  v.CreatedAt,
			Target:     lo.Ternary(ok, target[0], nil),
		}
	})

	return newsFeeds, nil
}
