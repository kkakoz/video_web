package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model"
)

type followRepo struct {
	ormx.IRepo[model.Follow]
}

var followOnce sync.Once
var _follow *followRepo

func Follow() *followRepo {
	followOnce.Do(func() {
		_follow = &followRepo{IRepo: ormx.NewRepo[model.Follow]()}
	})
	return _follow
}
