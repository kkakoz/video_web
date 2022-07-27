package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sync"
	"video_web/internal/consts"
	"video_web/internal/dto/request"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model"
	"video_web/internal/pkg/local"
	"video_web/pkg/errno"
)

type followLogic struct {
}

var followOnce sync.Once
var _follow *followLogic

func Follow() *followLogic {
	followOnce.Do(func() {
		_follow = &followLogic{}
	})
	return _follow
}

func (item *followLogic) Follow(ctx context.Context, req *request.FollowReq) (err error) {
	ctx, checkError := ormx.Begin(ctx)
	defer func() {
		err = checkError(err)
	}()
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	if req.GroupId == 0 {
		group, err := repo.FollowGroup().Get(ctx, opt.Where("user_id = ? and type = ?", user.ID, consts.FollowGroupTypeDefault))
		if err != nil {
			return err
		}
		req.GroupId = group.ID
	}
	exist, err := repo.FollowGroup().GetExist(ctx, opt.Where("user_id = ? and followed_user_id = ?", user.ID, req.FollowedUserId))
	if err != nil {
		return err
	}

	if req.Type == 1 { // 关注
		if exist { // 已经关注
			return errno.New400("已经关注该用户")
		} else { // 添加关注
			err = repo.Follow().Add(ctx, &model.Follow{
				UserId:         user.ID,
				FollowedUserId: req.FollowedUserId,
				GroupId:        req.GroupId,
			})
			if err != nil {
				return err
			}
		}
	} else { // 取消关注
		if !exist { // 没有关注关系
			return errno.New400("未关注该用户")
		} else { // 删除关注
			err = repo.Follow().Delete(ctx, opt.Where("user_id = ? and followed_user_id = ?", user.ID, req.FollowedUserId))
			if err != nil {
				return err
			}
		}
	}

	updateCount := lo.Ternary(req.Type == 1, 1, -1)

	err = repo.User().Updates(ctx, map[string]any{"fans_count": gorm.Expr("fans_count + ?", updateCount)},
		opt.Where("id = ?", req.FollowedUserId))
	if err != nil {
		return err
	}
	return repo.User().Updates(ctx, map[string]any{"follow_count": gorm.Expr("follow_count + ?", updateCount)},
		opt.Where("id = ?", user.ID))
}

func (item *followLogic) Fans(ctx context.Context, req *request.FollowFansReq) ([]*model.Follow, error) {
	list, err := repo.Follow().GetList(ctx, opt.NewOpts().Where("followed_user_id = ?", req.FollowedUserId).
		IsWhere(req.LastUserId != 0, "user_id < ?", req.LastUserId).Limit(consts.DefaultLimit).Order("id desc")...)
	if err != nil {
		return nil, err
	}

	userIds := lo.Map(list, func(follow *model.Follow, i int) int64 {
		return follow.UserId
	})
	userList, err := repo.User().GetList(ctx, opt.Where("id in ?", userIds))
	userGroup := lo.GroupBy(userList, func(user *model.User) int64 {
		return user.ID
	})

	lo.ForEach(list, func(follow *model.Follow, i int) {
		if users, ok := userGroup[follow.UserId]; ok {
			if len(users) >= 1 {
				follow.User = users[0]
			}
		}
	})
	return list, nil
}

func (item *followLogic) Followers(ctx context.Context, req *request.FollowersReq) ([]*model.Follow, error) {

	list, err := repo.Follow().GetList(ctx, opt.NewOpts().Where("user_id = ?", req.UserId).
		IsWhere(req.GroupId != 0, "group_id = ?", req.GroupId).
		IsWhere(req.LastUserId != 0, "followed_user_id < ?", req.LastUserId).Limit(consts.DefaultLimit).Order("id desc")...)
	if err != nil {
		return nil, err
	}

	userIds := lo.Map(list, func(follow *model.Follow, i int) int64 {
		return follow.FollowedUserId
	})
	userList, err := repo.User().GetList(ctx, opt.Where("id in ?", userIds))
	userGroup := lo.GroupBy(userList, func(user *model.User) int64 {
		return user.ID
	})

	lo.ForEach(list, func(follow *model.Follow, i int) {
		if users, ok := userGroup[follow.FollowedUserId]; ok {
			if len(users) >= 1 {
				follow.FollowedUser = users[0]
			}
		}
	})
	return list, nil
}

func (item *followLogic) GetFollowGroups(ctx context.Context) ([]*model.FollowGroup, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return repo.FollowGroup().GetList(ctx, opt.Where("user_id = ?", user.ID))
}

func (item *followLogic) IsFollow(ctx context.Context, req *request.FollowIsReq) (bool, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return false, err
	}
	return repo.Follow().GetExist(ctx, opt.Where("user_id = ? and followed_user_id = ?", user.ID, req.FollowedUserId))
}

func (item *followLogic) GroupAdd(ctx context.Context, req *request.FollowGroupAddReq) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	return repo.FollowGroup().Add(ctx, &model.FollowGroup{
		UserId:    user.ID,
		Type:      consts.FollowGroupTypeCustom,
		GroupName: req.Name,
	})
}

func (item *followLogic) Groups(ctx context.Context) ([]*model.FollowGroup, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return repo.FollowGroup().GetList(ctx, opt.Where("user_id = ?", user.ID))
}
