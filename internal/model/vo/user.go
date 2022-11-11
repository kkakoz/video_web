package vo

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
	LastLogin   int64          `json:"last_login"`
	Email       string         `json:"email" gorm:"type:varchar(255);uniqueIndex:email_index"`
	Phone       string         `json:"phone" gorm:"type:varchar(255);uniqueIndex:phone_index"`
	CreatedAt   timex.Time     `json:"created_at"`
	UpdatedAt   timex.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Followed    bool           `json:"followed"`
}
