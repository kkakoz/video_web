package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model"
)

type userSecurityRepo struct {
	ormx.IRepo[model.UserSecurity]
}

var userSecurityOnce sync.Once
var _userSecurity *userSecurityRepo

func UserSecurity() *userSecurityRepo {
	userSecurityOnce.Do(func() {
		_userSecurity = &userSecurityRepo{IRepo: ormx.NewRepo[model.UserSecurity]()}
	})
	return _userSecurity
}
