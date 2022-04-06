package model

type Follow struct {
	ID             int64 `json:"id"`
	UserId         int64 `json:"user_id"`
	FollowedUserId int64 `json:"followed_user_id"`
	GroupId        int64 `json:"group_id"`
	User           *User `json:"user"`
	FollowedUser   *User `json:"followed_user"`
	CreatedAt      int64 `gorm:"autoCreateTime" json:"createdAt"`
}

type FollowGroup struct {
	ID        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Type      uint8  `json:"type"`
	GroupName string `json:"group_name"`
}
