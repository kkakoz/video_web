package request

type FollowReq struct {
	FollowedUserId int64 `json:"followed_user_id"`
	GroupId        int64 `json:"group_id"`
	Type           uint8 `json:"type"`
}

type FollowFansReq struct {
	FollowedUserId int64 `query:"followed_user_id"`
	LastUserId     int64 `query:"last_user_id"`
}

type FollowersReq struct {
	UserId     int64 `query:"user_id"`
	GroupId    int64 `query:"group_id"`
	LastUserId int64 `query:"last_user_id"`
}

type FollowIsReq struct {
	FollowedUserId int64 `query:"followed_user_id"`
}

type FollowGroupAddReq struct {
	Name string `json:"name"`
}
