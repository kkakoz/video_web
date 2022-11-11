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
		Action:     req.Action,
	})
	return err
}

func (us newsfeed) UserNews(ctx context.Context, req *dto.UserNewsfeedList) ([]*vo.NewsFeed, error) {

	options := opt.NewOpts().Preload("User").Where("user_id = ? ", req.UserId).
		Order("created_at desc, id desc")
	if req.LastId != 0 {
		last, err := repo.Newsfeed().GetById(ctx, req.LastId)
		if err != nil {
			return nil, err
		}
		if last != nil {
			options = options.Where("created_at <= ? and id < ?", last.CreatedAt, last.ID)
		}
	}

	list, err := repo.Newsfeed().GetList(ctx, options...)
	if err != nil {
		return nil, err
	}
	return us.dealNewsFeed(ctx, list)
}

func (us newsfeed) List(ctx context.Context, req *dto.NewsfeedList) ([]*vo.NewsFeed, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	followerIds := make([]int64, 0)
	err = repo.Follow().Pluck(ctx, "id", followerIds, opt.NewOpts().Where("user_id = ?", user.ID)...)
	if err != nil {
		return nil, err
	}

	list, err := repo.Newsfeed().GetList(ctx, opt.NewOpts().Preload("User").Where("user_id in ?", followerIds).Order("created_at desc")...)

	return us.dealNewsFeed(ctx, list)
}

func (newsfeed) dealNewsFeed(ctx context.Context, list []*entity.Newsfeed) ([]*vo.NewsFeed, error) {
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

	newsFeeds := lo.Map(list, func(v *entity.Newsfeed, i int) *vo.NewsFeed {
		target, ok := videoMap[v.TargetId]
		if ok {
			return vo.ConvertToNewsFeed(v, target[0])
		} else {
			return vo.ConvertToNewsFeed(v, nil)
		}

	})

	return newsFeeds, nil
}

func (us newsfeed) FollowedNewsFeeds(ctx context.Context, req *dto.LastId) ([]*vo.NewsFeed, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	followedId := make([]int64, 0)
	err = repo.Follow().Pluck(ctx, "followed_user_id", &followedId, opt.NewOpts().Where("user_id = ?", user.ID)...)
	if err != nil {
		return nil, err
	}

	options := opt.NewOpts().Preload("User").Where("user_id in ?", append(followedId, user.ID)).Order("created_at desc, id desc")
	if req.LastId != 0 {
		last, err := repo.Newsfeed().GetById(ctx, req.LastId)
		if err != nil {
			return nil, err
		}
		if last != nil {
			options = options.Where("created_at <= ? and id < ?", last.CreatedAt, last.ID)
		}
	}

	list, err := repo.Newsfeed().GetList(ctx, options...)

	return us.dealNewsFeed(ctx, list)
}
