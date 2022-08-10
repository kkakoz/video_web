package entity

type Newsfeed struct {
	ID         int64              `json:"id"`
	UserId     int64              `json:"user_id"`
	TargetType NewsfeedTargetType `json:"target_type"`
	TargetId   uint8              `json:"target_id"`
	Content    string             `json:"content"`
}

type NewsfeedTargetType uint8

const (
	NewsfeedTargetTypeVideo NewsfeedTargetType = 1
)
