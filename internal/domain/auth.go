package domain

import (
	"context"
	"video_web/pkg/gormx"
)

type Auth struct {
	ID           int64  `json:"id"`
	IdentityType int32  `json:"identity_type"` // 登录类型
	Identifier   string `json:"identifier"`    // 标识
	Credential   string `json:"credential"`
	Salt         string `json:"salt"`

	UserId    int64 `json:"user_id"`
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

const (
	IdentityTypeEmail = iota + 1
)

type IAuthRepo interface {
	WithIdentifierAndType(identifier string, identityType int32) gormx.DBOption
	GetAuth(ctx context.Context, opts ...gormx.DBOption) (*Auth, error)
	DeleteAuth(ctx context.Context, id int64) error
	GetAuthByIdentify(ctx context.Context, identityType int32, identifier string) (*Auth, error)
}
