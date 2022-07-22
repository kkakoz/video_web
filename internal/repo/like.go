package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model"
)

type likeRepo struct {
	ormx.IRepo[model.Like]
}

var likeOnce sync.Once
var _like *likeRepo

func Like() *likeRepo {
	likeOnce.Do(func() {
		_like = &likeRepo{IRepo: ormx.NewRepo[model.Like]()}
	})
	return _like
}
