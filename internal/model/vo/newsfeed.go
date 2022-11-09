package vo

import (
	"video_web/internal/model/entity"
	"video_web/pkg/timex"
)

type NewsFeed struct {
	ID         int64                     `json:"id"`
	UserId     int64                     `json:"user_id"`
	TargetType entity.NewsfeedTargetType `json:"target_type"`
	TargetId   int64                     `json:"target_id"`
	Content    string                    `json:"content"`
	CreatedAt  timex.Time                `json:"created_at"`
	Target     any                       `json:"target"`
	Action     entity.NewsfeedAction     `json:"action"`

	User *entity.User `json:"user"`
}

func ConvertToNewsFeed(newsfeed *entity.Newsfeed, target any) *NewsFeed {
	return &NewsFeed{
		ID:         newsfeed.ID,
		UserId:     newsfeed.UserId,
		TargetType: newsfeed.TargetType,
		TargetId:   newsfeed.TargetId,
		Content:    newsfeed.Content,
		CreatedAt:  newsfeed.CreatedAt,
		Target:     target,
		User:       newsfeed.User,
		Action:     newsfeed.Action,
	}
}
