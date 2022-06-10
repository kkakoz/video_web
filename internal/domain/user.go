package domain

import (
	"context"
	"gorm.io/gorm"
	"video_web/internal/dto/request"
	"video_web/pkg/gormx"
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
	LastLogin   int64  `json:"last_login"`
	CreatedAt   int64  `gorm:"autoCreateTime"`
	UpdatedAt   int64  `gorm:"autoUpdateTime"`
	Auth        *Auth  `json:"auth"`
	DeletedAt   gorm.DeletedAt
}

type IUserLogic interface {
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUsers(ctx context.Context, ids []int64) ([]*User, error)
	Register(ctx context.Context, auth *request.RegisterReq) error
	Login(ctx context.Context, user *request.LoginReq) (string, error)
	GetCurUser(ctx context.Context, token string) (*User, error)
}

type IUserRepo interface {
	WithId(id int64) gormx.DBOption
	WithIds(ids []int64) gormx.DBOption
	Add(ctx context.Context, user *User) error
	Get(ctx context.Context, opts ...gormx.DBOption) (*User, error)
	GetList(ctx context.Context, opts ...gormx.DBOption) ([]*User, error)
}