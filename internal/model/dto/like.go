package dto

import "video_web/internal/model/entity"

type Like struct {
	TargetId   int64                 `json:"target_id"`
	TargetType entity.LikeTargetType `json:"target_type"`
	LikeType   bool                  `json:"like_type"`
}

type LikeIs struct {
	TargetId   int64                 `query:"target_id"`
	TargetType entity.LikeTargetType `query:"target_type"`
}
