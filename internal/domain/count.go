package domain

type Count struct {
	ID           int64 `json:"id"`
	TargetId     int64 `json:"target_id"`
	Type         int   `json:"type"`
	LikeCount    int   `json:"like_count"`
	CollectCount int   `json:"collect_count"`
	ViewCount    int   `json:"view_count"`
}
