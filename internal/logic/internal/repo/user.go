package repo

import (
	"sync"
	"video_web/internal/model/entity"

	"github.com/kkakoz/ormx"
)

type userRepo struct {
	ormx.IRepo[entity.User]
}

var userOnce sync.Once
var _user *userRepo

func User() *userRepo {
	userOnce.Do(func() {
		_user = &userRepo{IRepo: ormx.NewRepo[entity.User]()}
	})
	return _user
}
