package entity

import "video_web/pkg/timex"

type History struct {
	ID         int64      `json:"id"`
	UserId     int64      `json:"user_id" gorm:"uniqueIndex:user_video_index,priority:1"`
	VideoId    int64      `json:"video_id" gorm:"uniqueIndex:user_video_index,priority:2"`
	ResourceId int64      `json:"resource_id"`
	UpdatedAt  timex.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Duration   int64      `json:"duration"`
	IP         string     `json:"ip"`

	User  *User  `json:"user"`
	Video *Video `json:"video"`
}
