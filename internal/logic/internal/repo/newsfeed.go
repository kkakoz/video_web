package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type newsfeedRepo struct {
	ormx.IRepo[entity.Newsfeed]
}

var newsfeedOnce sync.Once
var _newsfeed *newsfeedRepo

func Newsfeed() *newsfeedRepo {
	newsfeedOnce.Do(func() {
		_newsfeed = &newsfeedRepo{IRepo: ormx.NewRepo[entity.Newsfeed]()}
	})
	return _newsfeed
}
