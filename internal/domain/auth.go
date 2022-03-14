package domain

import (
	"github.com/kkakoz/ormx"
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
	ormx.IRepo[Auth]
}
