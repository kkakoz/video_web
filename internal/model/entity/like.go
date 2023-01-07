package entity

import "video_web/pkg/timex"

type Like struct {
	UserId     int64      `json:"user_id"`
	TargetType TargetType `json:"target_type"`
	TargetId   int64      `json:"target_id"`
	Like       bool       `json:"like"`
	CreatedAt  timex.Time `json:"createdAt" gorm:"autoCreateTime" `
}

type TargetType uint8

const (
	TargetTypeVideo      = 1
	TargetTypeComment    = 2
	TargetTypeSubComment = 3
)
