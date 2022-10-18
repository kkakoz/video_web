package entity

import "video_web/pkg/timex"

type Collect struct {
	ID        int64      `json:"id"`
	UserId    int64      `json:"user_id"`
	TargetId  int64      `json:"target_id"`
	GroupId   int64      `json:"group_id"`
	CreatedAt timex.Time `json:"createdAt" gorm:"autoCreateTime" `

	User *User `json:"user"`
}

type CollectGroup struct {
	ID        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Type      uint8  `json:"type"`
	GroupName string `json:"group_name"`
}
