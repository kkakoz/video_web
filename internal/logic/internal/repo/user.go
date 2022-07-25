package repo

import (
	"sync"
	"video_web/internal/model"

	"github.com/kkakoz/ormx"
)

type userRepo struct {
	ormx.IRepo[model.User]
}

var userOnce sync.Once
var _user *userRepo

func User() *userRepo {
	userOnce.Do(func() {
		_user = &userRepo{IRepo: ormx.NewRepo[model.User]()}
	})
	return _user
}
