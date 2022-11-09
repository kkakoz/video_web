package dto

import "video_web/internal/model/entity"

type NewsfeedAdd struct {
	TargetType entity.NewsfeedTargetType `json:"target_type"`
	TargetId   int64                     `json:"target_id"`
	Content    string                    `json:"content"`
	Action     entity.NewsfeedAction     `json:"action"`
}

type NewsfeedList struct {
	LastId int64 `query:"last_id"`
}

type UserNewsfeedList struct {
	UserId int64 `query:"user_id"`
	LastId int64 `query:"last_id"`
}
