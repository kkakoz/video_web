package dto

type Follow struct {
	FollowedUserId int64 `json:"followed_user_id"`
	GroupId        int64 `json:"group_id"`
	Type           uint8 `json:"type"`
}

type FollowFans struct {
	FollowedUserId int64 `json:"followed_user_id"`
	LastUserId     int64 `json:"last_user_id"`
}

type Followers struct {
	UserId     int64 `json:"user_id"`
	GroupId    int64 `json:"group_id"`
	LastUserId int64 `json:"last_user_id"`
}

type FollowIs struct {
	FollowedUserId int64 `json:"followed_user_id"`
}

type FollowGroupAdd struct {
	Name string `json:"name"`
}
