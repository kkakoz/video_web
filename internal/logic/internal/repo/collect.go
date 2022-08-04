package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type collectRepo struct {
	ormx.IRepo[entity.Collect]
}

var collectOnce sync.Once
var _collect *collectRepo

func Collect() *collectRepo {
	collectOnce.Do(func() {
		_collect = &collectRepo{IRepo: ormx.NewRepo[entity.Collect]()}
	})
	return _collect
}
