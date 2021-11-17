package domain

import (
	"context"
	"github/kkakoz/video_web/internal/dto"
	"gorm.io/gorm"
)

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Brief       string `json:"brief"`
	FollowCount int64  `json:"follow_count"`
	FansCount   int64  `json:"fans_count"`
	LikeCount   int64  `json:"like_count"`
	State       int32  `json:"state"`
	Auth        *Auth  `json:"auth"`
	LastLogin   int64  `json:"last_login"`
	Created     int64  `gorm:"autoCreateTime"`
	Updated     int64  `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt
}

type IUserLogic interface {
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUsers(ctx context.Context, ids []int64) ([]*User, error)
	Register(ctx context.Context, auth *dto.RegisterReq) error
	Login(ctx context.Context, user *Auth) (string, error)
}

type IUserRepo interface {
	AddUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUserList(ctx context.Context, ids []int64) ([]*User, error)
}
