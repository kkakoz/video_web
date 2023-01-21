package entity

import "video_web/pkg/timex"

type Newsfeed struct {
	ID         int64              `json:"id"`
	UserId     int64              `json:"user_id"  gorm:"index:user_index"`
	TargetType NewsfeedTargetType `json:"target_type"`
	TargetId   int64              `json:"target_id"`
	Content    string             `json:"content"`
	CreatedAt  timex.Time         `json:"created_at"`
	Action     NewsfeedAction     `json:"action"`

	User *User `json:"user"`
}

type NewsfeedTargetType uint8

const (
	NewsfeedTargetTypeVideo NewsfeedTargetType = 1
)

type NewsfeedAction uint8

const (
	NewsfeedActionShare     NewsfeedAction = 1
	NewsfeedActionSend      NewsfeedAction = 2
	NewsfeedActionSendVideo NewsfeedAction = 3
)
