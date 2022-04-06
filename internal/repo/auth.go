package repo

import (
	"video_web/internal/model"

	"github.com/kkakoz/ormx"
)

type AuthRepo struct {
	ormx.IRepo[model.Auth]
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{
		ormx.NewRepo[model.Auth](),
	}
}
