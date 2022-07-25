package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model"
)

type danmuRepo struct {
	ormx.IRepo[model.Danmu]
}

var danmuOnce sync.Once
var _danmu *danmuRepo

func Danmu() *danmuRepo {
	danmuOnce.Do(func() {
		_danmu = &danmuRepo{IRepo: ormx.NewRepo[model.Danmu]()}
	})
	return _danmu
}
