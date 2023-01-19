package vo

import (
	"github.com/samber/lo"
	"video_web/internal/model/entity"
	"video_web/pkg/timex"
)

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

	//Like bool `json:"like"`
}

func ConvertToComment(comment *entity.Comment) *Comment {
	return &Comment{
		ID:           comment.ID,
		UserId:       comment.UserId,
		Username:     comment.Username,
		Avatar:       comment.Avatar,
		Content:      comment.Content,
		Top:          comment.Top,
		CommentCount: comment.CommentCount,
		LikeCount:    comment.LikeCount,
		CreatedAt:    comment.CreatedAt,
		UpdatedAt:    comment.UpdatedAt,
		SubComments: lo.Map(comment.SubComments, func(sub *entity.SubComment, i int) *SubComment {
			return ConvertToSubComment(sub)
		}),
	}
}

type SubComment struct {
	ID               int64      `json:"id"`
	CommentId        int64      `json:"comment_id" gorm:"index"`
	RootSubCommentId int64      `json:"root_sub_comment_id"` // 查看回复评论的 回复列表
	UserId           int64      `json:"user_id"`
	Username         string     `json:"username"`
	UserAvatar       string     `json:"user_avatar"`
	ToId             int64      `json:"to_id"`
	ToName           string     `json:"to_name"`
	Content          string     `json:"content"`
	CreatedAt        timex.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        timex.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Like bool `json:"like"`
}

func ConvertToSubComment(comments *entity.SubComment) *SubComment {
	return &SubComment{
		ID:               comments.ID,
		CommentId:        comments.CommentId,
		RootSubCommentId: comments.RootSubCommentId,
		UserId:           comments.UserId,
		Username:         comments.Username,
		UserAvatar:       comments.UserAvatar,
		ToId:             comments.ToId,
		ToName:           comments.ToName,
		Content:          comments.Content,
		CreatedAt:        comments.CreatedAt,
		UpdatedAt:        comments.UpdatedAt,
	}
}
