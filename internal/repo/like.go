package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/model"
)

type LikeRepo struct {
	ormx.IRepo[model.Like]
}

func NewLikeRepoRepo() *LikeRepo {
	return &LikeRepo{
		ormx.NewRepo[model.Like](),
	}
}
