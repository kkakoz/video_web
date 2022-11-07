package vo

import (
	"video_web/internal/model/entity"
	"video_web/pkg/timex"
)

type NewsFeedVo struct {
	ID         int64                     `json:"id"`
	UserId     int64                     `json:"user_id"`
	TargetType entity.NewsfeedTargetType `json:"target_type"`
	TargetId   int64                     `json:"target_id"`
	Content    string                    `json:"content"`
	CreatedAt  timex.Time                `json:"created_at"`
	Target     any                       `json:"target"`
}
