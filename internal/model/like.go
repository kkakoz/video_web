package model

type Like struct {
	UserId     int64 `json:"user_id"`
	TargetType uint8 `json:"target_type"`
	TargetId   int64 `json:"target_id"`
}
