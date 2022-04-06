package repo

import (
	"video_web/internal/model"

	"github.com/kkakoz/ormx"
)

type UserRepo struct {
	ormx.IRepo[model.User]
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		ormx.NewRepo[model.User](),
	}
}
