package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/model"
)

type SubCommentRepo struct {
	ormx.IRepo[model.SubComment]
}

func NewSubCommentRepo() *SubCommentRepo {
	return &SubCommentRepo{IRepo: ormx.NewRepo[model.SubComment]()}
}
