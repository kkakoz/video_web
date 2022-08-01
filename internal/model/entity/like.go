package entity

import "video_web/pkg/timex"

type Like struct {
	UserId     int64          `json:"user_id"`
	TargetType LikeTargetType `json:"target_type"`
	TargetId   int64          `json:"target_id"`
	CreatedAt  timex.Time     `json:"createdAt" gorm:"autoCreateTime" `
}

type LikeTargetType uint8

const (
	LikeTargetTypeVideo      = 1
	LikeTargetTypeCollection = 2
)
