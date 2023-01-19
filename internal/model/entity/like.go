package entity

import "video_web/pkg/timex"

type Like struct {
	UserId     int64      `json:"user_id" gorm:"uniqueIndex:target_index,priority:3"`
	TargetType TargetType `json:"target_type" gorm:"uniqueIndex:target_index,priority:2"`
	TargetId   int64      `json:"target_id" gorm:"uniqueIndex:target_index,priority:1"`
	Like       LikeState  `json:"like"`
	CreatedAt  timex.Time `json:"createdAt" gorm:"autoCreateTime" `
}

type TargetType uint8

const (
	TargetTypeVideo      TargetType = 1
	TargetTypeComment    TargetType = 2
	TargetTypeSubComment TargetType = 3
)

type LikeState uint8

const (
	LikeStateCancel LikeState = 0
	LikeStateLike   LikeState = 1
	//LikeStateDisLike LikeState = 2
)
