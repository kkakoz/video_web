package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type historyRepo struct {
	ormx.IRepo[entity.History]
}

var historyOnce sync.Once
var _history *historyRepo

func History() *historyRepo {
	historyOnce.Do(func() {
		_history = &historyRepo{IRepo: ormx.NewRepo[entity.History]()}
	})
	return _history
}
