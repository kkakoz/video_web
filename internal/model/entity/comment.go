package entity

import "video_web/pkg/timex"

type Comment struct {
	ID           int64         `json:"id"`
	TargetType   uint8         `json:"target_type" gorm:"index:target_index,priority:2"`
	TargetId     int64         `json:"target_id" gorm:"index:target_index,priority:1"`
	UserId       int64         `json:"user_id"`
	Username     string        `json:"username"`
	Avatar       string        `json:"avatar"`
	Content      string        `json:"content"`
	Top          bool          `json:"top"` // 置顶
	CommentCount int64         `json:"comment_count"`
	LikeCount    int64         `json:"like_count"`
	CreatedAt    timex.Time    `json:"created_at"`
	UpdatedAt    timex.Time    `json:"updated_at"`
	SubComments  []*SubComment `json:"sub_comments"`
}

type CommentTargetType uint8

const (
	CommentTargetTypeVideo CommentTargetType = 1
)

type SubComment struct {
	ID               int64      `json:"id"`
	CommentId        int64      `json:"comment_id" gorm:"index"`
	RootSubCommentId int64      `json:"root_sub_comment_id"` // 查看回复评论的 回复列表
	FromId           int64      `json:"from_id"`
	FromName         string     `json:"from_name"`
	FromAvatar       string     `json:"from_avatar"`
	ToId             int64      `json:"to_id"`
	ToName           string     `json:"to_name"`
	ParentId         int64      `json:"parent_id"`
	Content          string     `json:"content"`
	CreatedAt        timex.Time `json:"created_at"`
	UpdatedAt        timex.Time `json:"updated_at"`
}
