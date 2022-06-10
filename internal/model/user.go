package model

import (
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
	LastLogin   int64  `json:"last_login"`
	Email       string `json:"email" gorm:"type:varchar(255);uniqueIndex:email_index"`
	Phone       string `json:"phone" gorm:"type:varchar(255);uniqueIndex:phone_index"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime"`
	// Auth        *Auth          `json:"auth"`
	UserSecurity *UserSecurity  `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type UserSecurity struct {
	UserId   int64  `json:"user_id"`
	Password string `json:"password"`
	Salt     string `json:"salt" gorm:"type:varchar(255)"`
}
