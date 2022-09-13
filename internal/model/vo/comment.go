package vo

import "video_web/pkg/timex"

type Comment struct {
	ID           int64         `json:"id"`
	UserId       int64         `json:"user_id"`
	Username     string        `json:"username"`
	Avatar       string        `json:"avatar"`
	Content      string        `json:"content"`
	Top          bool          `json:"top"`           // 置顶
	CommentCount int64         `json:"comment_count"` // 子评论数量
	LikeCount    int64         `json:"like_count"`
	CreatedAt    timex.Time    `json:"created_at"`
	UpdatedAt    timex.Time    `json:"updated_at"`
	SubComments  []*SubComment `json:"sub_comments"`

	Like bool `json:"like"`
}

type SubComment struct {
	ID               int64      `json:"id"`
	CommentId        int64      `json:"comment_id" gorm:"index"`
	RootSubCommentId int64      `json:"root_sub_comment_id"` // 查看回复评论的 回复列表
	FromId           int64      `json:"from_id"`
	FromName         string     `json:"from_name"`
	FromAvatar       string     `json:"from_avatar"`
	ToId             int64      `json:"to_id"`
	ToName           string     `json:"to_name"`
	Content          string     `json:"content"`
	CreatedAt        timex.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        timex.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Like bool `json:"like"`
}
