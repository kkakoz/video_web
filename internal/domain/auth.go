package domain

import (
	"context"
)

type Auth struct {
	ID           int64  `json:"id"`
	IdentityType int32  `json:"identity_type"` // 登录类型
	Identifier   string `json:"identifier"`    // 标识
	Credential   string `json:"credential"`
	UserId       int64  `json:"user_id"`
	Created      int64  `gorm:"autoCreateTime"`
	Updated      int64  `gorm:"autoUpdateTime"`
}

const (
	IdentityTypeEmail = iota + 1
)

type IAuthRepo interface {
	GetAuth(ctx context.Context, id int64) (*Auth, error)
	DeleteAuth(ctx context.Context, id int64) error
	GetAuthByEmail(ctx context.Context, email string) (*Auth, error)
}
