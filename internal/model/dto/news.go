package dto

type NewsfeedAdd struct {
	TargetType int64  `json:"target_type"`
	TargetId   int64  `json:"target_id"`
	Content    string `json:"content"`
}
