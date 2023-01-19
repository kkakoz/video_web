package dto

import "video_web/internal/model/entity"

type Like struct {
	TargetId   int64             `json:"target_id"`
	TargetType entity.TargetType `json:"target_type"`
	LikeState  entity.LikeState  `json:"like_state"` // 0 取消 1 like 2 dislike
}

type UserLikeList struct {
	UserId     int64             `json:"user_id"`
	TargetIds  []int64           `json:"target_ids"`
	TargetType entity.TargetType `json:"target_type"`
}

type IsLike struct {
	UserId     int64             `json:"user_id"`
	TargetType entity.TargetType `json:"target_type"`
	TargetId   int64             `json:"target_id"`
}
