package entity

import (
	"gorm.io/gorm"
	"video_web/pkg/timex"
)

type User struct {
	ID          int64          `json:"id"`
	Type        UserType       `json:"type" gorm:"default:1"`
	Name        string         `json:"name"`
	Avatar      string         `json:"avatar"`
	Brief       string         `json:"brief"`
	FollowCount int64          `json:"follow_count"`
	FansCount   int64          `json:"fans_count"`
	LikeCount   int64          `json:"like_count"`
	State       UserState      `json:"state" gorm:"default:1"`
	LastLogin   timex.Time     `json:"last_login"`
	Email       string         `json:"email" gorm:"type:varchar(255);uniqueIndex:email_index"`
	Phone       string         `json:"phone" gorm:"type:varchar(255);index:phone_index"`
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
	UserId   int64  `json:"user_id" gorm:"index:user_index"`
	Password string `json:"password"`
	Salt     string `json:"salt" gorm:"type:varchar(255)"`
}
