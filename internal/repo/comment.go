package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/domain"
)

var _ domain.ICommentRepo = (*CommentRepo)(nil)

type CommentRepo struct {
	ormx.IRepo[domain.Comment]
}

func NewCommentRepo() domain.ICommentRepo {
	return &CommentRepo{IRepo: ormx.NewRepo[domain.Comment]()}
}
