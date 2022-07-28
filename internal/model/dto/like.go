package dto

type Like struct {
	TargetId   int64 `json:"target_id"`
	TargetType uint8 `json:"target_type"`
	LikeType   bool  `json:"like_type"`
}

type LikeIs struct {
	TargetId   int64 `query:"target_id"`
	TargetType uint8 `query:"target_type"`
}
