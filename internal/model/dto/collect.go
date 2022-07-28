package dto

type CollectAdd struct {
	TargetType uint8 `json:"target_type"`
	TargetId   int64 `json:"target_id"`
	GroupId    int64 `json:"group_id"`
}

type CollectGroupAdd struct {
	ID     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
}

type CollectIs struct {
	TargetType uint8 `query:"target_type"`
	TargetId   int64 `query:"target_id"`
}

type Collects struct {
	GroupId int64 `query:"group_id"`
}
