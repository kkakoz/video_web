package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type danmuRepo struct {
	ormx.IRepo[entity.Danmu]
}

var danmuOnce sync.Once
var _danmu *danmuRepo

func Danmu() *danmuRepo {
	danmuOnce.Do(func() {
		_danmu = &danmuRepo{IRepo: ormx.NewRepo[entity.Danmu]()}
	})
	return _danmu
}
