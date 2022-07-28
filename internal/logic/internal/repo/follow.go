package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type followRepo struct {
	ormx.IRepo[entity.Follow]
}

var followOnce sync.Once
var _follow *followRepo

func Follow() *followRepo {
	followOnce.Do(func() {
		_follow = &followRepo{IRepo: ormx.NewRepo[entity.Follow]()}
	})
	return _follow
}
