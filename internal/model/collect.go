package model

import "video_web/pkg/timex"

type Collect struct {
	ID         int64      `json:"id"`
	UserId     int64      `json:"user_id"`
	TargetType uint8      `json:"target_type"`
	TargetId   int64      `json:"target_id"`
	GroupId    int64      `json:"group_id"`
	CreatedAt  timex.Time `gorm:"autoCreateTime" json:"createdAt"`

	User *User `json:"user"`
}

type CollectGroup struct {
	ID        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Type      uint8  `json:"type"`
	GroupName string `json:"group_name"`
}
