package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model"
)

type followGroupRepo struct {
	ormx.IRepo[model.FollowGroup]
}

var followGroupOnce sync.Once
var _followGroup *followGroupRepo

func FollowGroup() *followGroupRepo {
	followGroupOnce.Do(func() {
		_followGroup = &followGroupRepo{IRepo: ormx.NewRepo[model.FollowGroup]()}
	})
	return _followGroup
}
