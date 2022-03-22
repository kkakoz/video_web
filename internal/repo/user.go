package repo

import (
	"video_web/internal/domain"

	"github.com/kkakoz/ormx"
)

var _ domain.IUserRepo = (*UserRepo)(nil)

type UserRepo struct {
	ormx.IRepo[domain.User]
}

func NewUserRepo() domain.IUserRepo {
	return &UserRepo{
		ormx.NewRepo[domain.User](),
	}
}
