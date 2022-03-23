package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/domain"
)

var _ domain.ILikeRepo = (*LikeRepo)(nil)

type LikeRepo struct {
	ormx.IRepo[domain.Like]
}

func NewLikeRepoRepo() domain.ILikeRepo {
	return &LikeRepo{
		ormx.NewRepo[domain.Like](),
	}
}
