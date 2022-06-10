package request

type LikeReq struct {
	TargetId   int64 `json:"target_id"`
	TargetType uint8 `json:"target_type"`
	LikeType   bool  `json:"like_type"`
}

type LikeIsReq struct {
	TargetId   int64 `query:"target_id"`
	TargetType uint8 `query:"target_type"`
}
