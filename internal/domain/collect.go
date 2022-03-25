package domain

import (
	"context"
	"github.com/kkakoz/ormx"
	"video_web/internal/dto/request"
)

type Collect struct {
	ID           int64 `json:"id"`
	UserId       int64 `json:"user_id"`
	TargetType   uint8 `json:"target_type"`
	TargetId     int64 `json:"target_id"`
	GroupId      int64 `json:"group_id"`
	User         *User `json:"user"`
	FollowedUser *User `json:"followed_user"`
	CreatedAt    int64 `gorm:"autoCreateTime" json:"createdAt"`
}

type CollectGroup struct {
	ID        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Type      uint8  `json:"type"`
	GroupName string `json:"group_name"`
}

type ICollectLogic interface {
	Collect(ctx context.Context, req *request.CollectAddReq) error
	Collects(ctx context.Context, req *request.CollectsReq) ([]*Follow, error)
	CollectGroupAdd(ctx context.Context, req *request.CollectGroupAddReq) error
	IsCollect(ctx context.Context, req *request.CollectIsReq) (bool, error)
}

type ICollectRepo interface {
	ormx.IRepo[Follow]
}

type ICollectGroupRepo interface {
	ormx.IRepo[FollowGroup]
}
