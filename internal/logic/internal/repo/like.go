package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type likeRepo struct {
	ormx.IRepo[entity.Like]
}

var likeOnce sync.Once
var _like *likeRepo

func Like() *likeRepo {
	likeOnce.Do(func() {
		_like = &likeRepo{IRepo: ormx.NewRepo[entity.Like]()}
	})
	return _like
}
