package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/model"
)

type FollowRepo struct {
	ormx.IRepo[model.Follow]
}

func NewFollowRepo() *FollowRepo {
	return &FollowRepo{
		ormx.NewRepo[model.Follow](),
	}
}
