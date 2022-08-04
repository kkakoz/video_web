package entity

import "video_web/pkg/timex"

type History struct {
	ID        int64      `json:"id"`
	UserId    int64      `json:"user_id"`
	VideoId   int64      `json:"video_id"`
	CreatedAt timex.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt timex.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Duration  int64      `json:"duration"`
	IP        string     `json:"ip"`

	User  *User  `json:"user"`
	Video *Video `json:"video"`
}
