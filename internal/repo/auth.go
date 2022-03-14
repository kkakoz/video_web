package repo

import (
	"video_web/internal/domain"

	"github.com/kkakoz/ormx"
)

var _ domain.IAuthRepo = (*AuthRepo)(nil)

type AuthRepo struct {
	ormx.Repo[domain.Auth]
}

func NewAuthRepo() domain.IAuthRepo {
	return &AuthRepo{}
}
