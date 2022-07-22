package model

import "video_web/pkg/timex"

type Like struct {
	UserId     int64      `json:"user_id"`
	TargetType uint8      `json:"target_type"`
	TargetId   int64      `json:"target_id"`
	CreatedAt  timex.Time `gorm:"autoCreateTime" json:"createdAt"`
}
