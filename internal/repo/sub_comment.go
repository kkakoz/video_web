package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/domain"
)

var _ domain.ISubCommentRepo = (*SubCommentRepo)(nil)

type SubCommentRepo struct {
	ormx.IRepo[domain.SubComment]
}

func NewSubCommentRepo() domain.ISubCommentRepo {
	return &SubCommentRepo{IRepo: ormx.NewRepo[domain.SubComment]()}
}
