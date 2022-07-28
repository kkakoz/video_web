package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type commentRepo struct {
	ormx.IRepo[entity.Comment]
}

var commentOnce sync.Once
var _comment *commentRepo

func Comment() *commentRepo {
	commentOnce.Do(func() {
		_comment = &commentRepo{IRepo: ormx.NewRepo[entity.Comment]()}
	})
	return _comment
}
