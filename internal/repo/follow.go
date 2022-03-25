package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/domain"
)

var _ domain.IFollowRepo = (*FollowRepo)(nil)

type FollowRepo struct {
	ormx.IRepo[domain.Follow]
}

func NewFollowRepo() domain.IFollowRepo {
	return &FollowRepo{
		ormx.NewRepo[domain.Follow](),
	}
}
