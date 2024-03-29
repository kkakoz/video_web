package logic

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/ormx/opt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sync"
	"video_web/internal/consts"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
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

func (followLogic) Follow(ctx context.Context, req *dto.Follow) (err error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	return ormx.Transaction(ctx, func(ctx context.Context) error {
		if req.GroupId == 0 {
			group, err := repo.FollowGroup().Get(ctx, opt.Where("user_id = ? and type = ?", user.ID, entity.FollowGroupTypeNormal))
			if err != nil {
				return err
			}
			if group == nil { // 没找到默认关注分组
				group = &entity.FollowGroup{
					UserId:    user.ID,
					Type:      entity.FollowGroupTypeNormal,
					GroupName: "默认关注",
				}
				err = repo.FollowGroup().Add(ctx, group)
				if err != nil {
					return err
				}
			}
			req.GroupId = group.ID
		}
		exist, err := repo.Follow().GetExist(ctx, opt.Where("user_id = ? and followed_user_id = ?", user.ID, req.FollowedUserId))
		if err != nil {
			return err
		}

		updateCount := 0

		if req.Type == 1 { // 关注
			if exist { // 已经关注
				return errno.New400("已经关注该用户")
			} else { // 添加关注
				err = repo.Follow().Add(ctx, &entity.Follow{
					UserId:         user.ID,
					FollowedUserId: req.FollowedUserId,
					GroupId:        req.GroupId,
				})
				if err != nil {
					return err
				}
				updateCount += 1
			}
		} else { // 取消关注
			if !exist { // 没有关注关系
				return nil
			} else { // 删除关注
				err = repo.Follow().Delete(ctx, opt.Where("user_id = ? and followed_user_id = ?", user.ID, req.FollowedUserId))
				if err != nil {
					return err
				}
				updateCount -= 1
			}
		}

		err = repo.User().Updates(ctx, map[string]any{"fans_count": gorm.Expr("fans_count + ?", updateCount)},
			opt.Where("id = ?", req.FollowedUserId))
		if err != nil {
			return err
		}
		return repo.User().Updates(ctx, map[string]any{"follow_count": gorm.Expr("follow_count + ?", updateCount)},
			opt.Where("id = ?", user.ID))
	})
}

// Fans 粉丝列表
func (followLogic) Fans(ctx context.Context, req *dto.FollowFans) ([]*entity.Follow, error) {
	last, err := repo.Follow().GetById(ctx, req.LastId)
	if err != nil {
		return nil, err
	}

	options := opt.NewOpts().Where("followed_user_id = ?", req.UserId).Limit(consts.DefaultLimit).Order("id desc")
	if req.LastId != 0 {
		options = options.Where("created_at <= ? and id < ?", last.CreatedAt, req.LastId)
	}
	list, err := repo.Follow().GetList(ctx, options...)
	if err != nil {
		return nil, err
	}

	userIds := lo.Map(list, func(follow *entity.Follow, i int) int64 {
		return follow.UserId
	})
	userList, err := repo.User().GetList(ctx, opt.Where("id in ?", userIds))
	userGroup := lo.GroupBy(userList, func(user *entity.User) int64 {
		return user.ID
	})

	lo.ForEach(list, func(follow *entity.Follow, i int) {
		if users, ok := userGroup[follow.UserId]; ok {
			if len(users) >= 1 {
				follow.User = users[0]
			}
		}
	})
	return list, nil
}

// Followers 关注列表
func (followLogic) Followers(ctx context.Context, req *dto.Followers) ([]*entity.User, error) {

	list, err := repo.Follow().GetList(ctx, opt.NewOpts().Where("user_id = ?", req.UserId).
		IsWhere(req.GroupId != 0, "group_id = ?", req.GroupId).
		IsWhere(req.LastId != 0, "id < ?", req.LastId).Limit(consts.DefaultLimit).
		Order("id desc").Preload("FollowedUser")...)

	res := lo.Map(list, func(value *entity.Follow, i int) *entity.User {
		return value.FollowedUser
	})

	return res, err
}

func (followLogic) GetFollowGroups(ctx context.Context) ([]*entity.FollowGroup, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return repo.FollowGroup().GetList(ctx, opt.Where("user_id = ?", user.ID))
}

func (followLogic) IsFollow(ctx context.Context, req *dto.FollowIs) (bool, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return false, err
	}
	return repo.Follow().GetExist(ctx, opt.Where("user_id = ? and followed_user_id = ?", user.ID, req.FollowedUserId))
}

func (followLogic) GroupAdd(ctx context.Context, req *dto.FollowGroupAdd) error {
	user, err := local.GetUser(ctx)
	if err != nil {
		return err
	}
	return repo.FollowGroup().Add(ctx, &entity.FollowGroup{
		UserId:    user.ID,
		Type:      entity.FollowGroupTypeCustom,
		GroupName: req.Name,
	})
}

func (followLogic) Groups(ctx context.Context) ([]*entity.FollowGroup, error) {
	user, err := local.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return repo.FollowGroup().GetList(ctx, opt.Where("user_id = ?", user.ID))
}

func (l followLogic) FollowedUser(ctx context.Context, userId int64) bool {
	current, exist := local.GetUserExist(ctx)
	if exist {
		followed, err := repo.Follow().GetExist(ctx, opt.Where("user_id = ? and followed_user_id = ?", current.ID, userId))
		if err != nil {
			return false
		}
		return followed
	}
	return false
}
