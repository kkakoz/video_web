package entity

import (
	"gorm.io/gorm"
	"video_web/pkg/timex"
)

type User struct {
	ID          int64          `json:"id"`
	Type        uint8          `json:"type"`
	Name        string         `json:"name"`
	Avatar      string         `json:"avatar"`
	Brief       string         `json:"brief"`
	FollowCount int64          `json:"follow_count"`
	FansCount   int64          `json:"fans_count"`
	LikeCount   int64          `json:"like_count"`
	State       int32          `json:"state"`
	LastLogin   timex.Time     `json:"last_login"`
	Email       string         `json:"email" gorm:"type:varchar(255);uniqueIndex:email_index"`
	Phone       string         `json:"phone" gorm:"type:varchar(255);uniqueIndex:phone_index"`
	CreatedAt   timex.Time     `json:"created_at"`
	UpdatedAt   timex.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`

	UserSecurity *UserSecurity `json:"-"`
}

type UserType uint8

const (
	UserTypeNormal UserType = 1
	UserTypeAdmin  UserType = 2
)

type UserState uint8

const (
	UserStateRegister = 0
	UserStateActive   = 1
)

type UserSecurity struct {
	UserId   int64  `json:"user_id"`
	Password string `json:"password"`
	Salt     string `json:"salt" gorm:"type:varchar(255)"`
}
