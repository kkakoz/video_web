package entity

import "video_web/pkg/timex"

type Follow struct {
	ID             int64      `json:"id"`
	UserId         int64      `json:"user_id" gorm:"index:user_index"`
	FollowedUserId int64      `json:"followed_user_id" gorm:"index:followed_user_index"`
	GroupId        int64      `json:"group_id"`
	CreatedAt      timex.Time `gorm:"autoCreateTime" json:"createdAt"`

	User         *User `json:"user"`
	FollowedUser *User `json:"followed_user"`
}

type FollowGroup struct {
	ID        int64           `json:"id"`
	UserId    int64           `json:"user_id"`
	Type      FollowGroupType `json:"type"`
	GroupName string          `json:"group_name"`
}

type FollowGroupType uint8

const (
	FollowGroupTypeNormal  FollowGroupType = 1
	FollowGroupTypeSpecial FollowGroupType = 2
	FollowGroupTypeCustom  FollowGroupType = 3
)
