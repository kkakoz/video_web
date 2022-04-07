package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/model"
)

type CommentRepo struct {
	ormx.IRepo[model.Comment]
}

func NewCommentRepo() *CommentRepo {
	return &CommentRepo{IRepo: ormx.NewRepo[model.Comment]()}
}
