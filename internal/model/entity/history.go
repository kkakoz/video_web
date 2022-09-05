package entity

import "video_web/pkg/timex"

type History struct {
	ID         int64       `json:"id"`
	UserId     int64       `json:"user_id"`
	TargetType HistoryType `json:"target_type"`
	TargetId   int64       `json:"target_id"`
	UpdatedAt  timex.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	Duration   int64       `json:"duration"`
	IP         string      `json:"ip"`

	User     *User     `json:"user"`
	Resource *Resource `json:"resource"`
}

type HistoryType uint8

const (
	HistoryTypeVideo HistoryType = 1 // 视频
)
