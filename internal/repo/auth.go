package repo

import (
	"video_web/internal/domain"

	"github.com/kkakoz/ormx"
)

var _ domain.IAuthRepo = (*AuthRepo)(nil)

type AuthRepo struct {
	ormx.IRepo[domain.Auth]
}

func NewAuthRepo() domain.IAuthRepo {
	return &AuthRepo{
		ormx.NewRepo[domain.Auth](),
	}
}
