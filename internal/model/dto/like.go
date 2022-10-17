package dto

import "video_web/internal/model/entity"

type Like struct {
	TargetId   int64                 `json:"target_id"`
	TargetType entity.LikeTargetType `json:"target_type"`
	LikeType   bool                  `json:"like_type"`
	Like       bool                  `json:"like"`
}

type LikeIs struct {
	UserId     int64                 `json:"user_id"`
	TargetIds  []int64               `json:"target_id"`
	TargetType entity.LikeTargetType `json:"target_type"`
}
