package dto

type Follow struct {
	FollowedUserId int64 `json:"followed_user_id"`
	GroupId        int64 `json:"group_id"`
	Type           uint8 `json:"type"` // 1 关注 2 取关
}

type FollowFans struct {
	UserId int64 `query:"user_id"`
	LastId int64 `query:"last_id"`
}

type Followers struct {
	UserId  int64 `query:"user_id" binding:"required"`
	GroupId int64 `query:"group_id"`
	LastId  int64 `query:"last_id"`
}

type FollowIs struct {
	FollowedUserId int64 `query:"followed_user_id"`
}

type FollowGroupAdd struct {
	Name string `json:"name"`
}
