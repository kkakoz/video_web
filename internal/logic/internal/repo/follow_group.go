package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type followGroupRepo struct {
	ormx.IRepo[entity.FollowGroup]
}

var followGroupOnce sync.Once
var _followGroup *followGroupRepo

func FollowGroup() *followGroupRepo {
	followGroupOnce.Do(func() {
		_followGroup = &followGroupRepo{IRepo: ormx.NewRepo[entity.FollowGroup]()}
	})
	return _followGroup
}
