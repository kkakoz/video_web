package domain

import (
	"context"
	"github.com/kkakoz/ormx"
	"video_web/internal/dto/request"
)

type Follow struct {
	ID             int64 `json:"id"`
	UserId         int64 `json:"user_id"`
	FollowedUserId int64 `json:"followed_user_id"`
	GroupId        int64 `json:"group_id"`
	User           *User `json:"user"`
	FollowedUser   *User `json:"followed_user"`
	CreatedAt      int64 `gorm:"autoCreateTime" json:"createdAt"`
}

type FollowGroup struct {
	ID        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Type      uint8  `json:"type"`
	GroupName string `json:"group_name"`
}

type IFollowLogic interface {
	Follow(ctx context.Context, req *request.FollowReq) error
	Fans(ctx context.Context, req *request.FollowFansReq) ([]*Follow, error)
	Followers(ctx context.Context, req *request.FollowersReq) ([]*Follow, error)
	IsFollow(ctx context.Context, req *request.FollowIsReq) (bool, error)
	GroupAdd(ctx context.Context, req *request.FollowGroupAddReq) error
	Groups(ctx context.Context) ([]*FollowGroup, error)
}

type IFollowRepo interface {
	ormx.IRepo[Follow]
}

type IFollowGroupRepo interface {
	ormx.IRepo[FollowGroup]
}
