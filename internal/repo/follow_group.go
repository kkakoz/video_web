package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/model"
)

type FollowGroupRepo struct {
	ormx.IRepo[model.FollowGroup]
}

func NewFollowGroupRepo() *FollowGroupRepo {
	return &FollowGroupRepo{
		ormx.NewRepo[model.FollowGroup](),
	}
}
