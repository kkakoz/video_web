package entity

type Notice struct {
	ID         int64      `json:"id"`
	Content    string     `json:"content"`
	FromUserId int64      `json:"from_user_id"`
	CreatedAt  int64      `json:"created_at" gorm:"autoCreateTime"`
	IsRead     bool       `json:"is_read"` // 是否看过
	UserId     int64      `json:"user_id" gorm:"index:user_index"`
	TargetType TargetType `json:"target_type"`
	TargetId   int64      `json:"target_id"`

	FromUser *User `json:"from_user"`
}
