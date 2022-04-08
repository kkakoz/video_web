package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Avatar      string         `json:"avatar"`
	Brief       string         `json:"brief"`
	FollowCount int64          `json:"follow_count"`
	FansCount   int64          `json:"fans_count"`
	LikeCount   int64          `json:"like_count"`
	State       int32          `json:"state"`
	LastLogin   int64          `json:"last_login"`
	CreatedAt   int64          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64          `json:"updated_at" gorm:"autoUpdateTime"`
	Auth        *Auth          `json:"auth"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
