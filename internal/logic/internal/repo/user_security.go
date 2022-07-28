package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type userSecurityRepo struct {
	ormx.IRepo[entity.UserSecurity]
}

var userSecurityOnce sync.Once
var _userSecurity *userSecurityRepo

func UserSecurity() *userSecurityRepo {
	userSecurityOnce.Do(func() {
		_userSecurity = &userSecurityRepo{IRepo: ormx.NewRepo[entity.UserSecurity]()}
	})
	return _userSecurity
}
