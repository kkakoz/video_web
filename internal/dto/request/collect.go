package request

type CollectAddReq struct {
	TargetType uint8 `json:"target_type"`
	TargetId   int64 `json:"target_id"`
	GroupId    int64 `json:"group_id"`
}

type CollectGroupAddReq struct {
	ID     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
}

type CollectIsReq struct {
	TargetType uint8 `query:"target_type"`
	TargetId   int64 `query:"target_id"`
}

type CollectsReq struct {
	GroupId int64 `query:"group_id"`
}
