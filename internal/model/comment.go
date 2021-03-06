package model

type Comment struct {
	ID           int64         `json:"id"`
	TargetType   uint8         `json:"target_type" gorm:"index:target_index,priority:2"`
	TargetId     int64         `json:"target_id" gorm:"index:target_index,priority:1"`
	UserId       int64         `json:"user_id"`
	Username     string        `json:"username"`
	Avatar       string        `json:"avatar"`
	Content      string        `json:"content"`
	CommentCount int64         `json:"comment_count"`
	LikeCount    int64         `json:"like_count"`
	SubComments  []*SubComment `json:"sub_comments"`
}

type SubComment struct {
	ID         int64  `json:"id"`
	CommentId  int64  `json:"comment_id" gorm:"index"`
	FromId     int64  `json:"from_id"`
	FromName   string `json:"from_name"`
	FromAvatar string `json:"from_avatar"`
	ToId       int64  `json:"to_id"`
	ToName     string `json:"to_name"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
}
