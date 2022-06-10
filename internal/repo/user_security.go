package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/model"
)

type UserSecurityRepo struct {
	ormx.IRepo[model.UserSecurity]
}

func NewUserSecurityRepo() *UserSecurityRepo {
	return &UserSecurityRepo{
		ormx.NewRepo[model.UserSecurity](),
	}
}
