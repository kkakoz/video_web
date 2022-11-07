package dto

import "video_web/internal/model/entity"

type NewsfeedAdd struct {
	TargetType entity.NewsfeedTargetType `json:"target_type"`
	TargetId   int64                     `json:"target_id"`
	Content    string                    `json:"content"`
}

type NewsfeedList struct {
}
