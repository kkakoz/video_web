package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/domain"
)

var _ domain.IFollowGroupRepo = (*FollowGroupRepo)(nil)

type FollowGroupRepo struct {
	ormx.IRepo[domain.FollowGroup]
}

func NewFollowGroupRepo() domain.IFollowGroupRepo {
	return &FollowGroupRepo{
		ormx.NewRepo[domain.FollowGroup](),
	}
}
