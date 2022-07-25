package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model"
)

type commentRepo struct {
	ormx.IRepo[model.Comment]
}

var commentOnce sync.Once
var _comment *commentRepo

func Comment() *commentRepo {
	commentOnce.Do(func() {
		_comment = &commentRepo{IRepo: ormx.NewRepo[model.Comment]()}
	})
	return _comment
}
