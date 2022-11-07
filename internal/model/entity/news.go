package entity

import "video_web/pkg/timex"

type Newsfeed struct {
	ID         int64              `json:"id"`
	UserId     int64              `json:"user_id"`
	TargetType NewsfeedTargetType `json:"target_type"`
	TargetId   int64              `json:"target_id"`
	Content    string             `json:"content"`
	CreatedAt  timex.Time         `json:"created_at"`
}

type NewsfeedTargetType uint8

const (
	NewsfeedTargetTypeVideo NewsfeedTargetType = 1
)
