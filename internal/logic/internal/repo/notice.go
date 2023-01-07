package repo

import (
	"github.com/kkakoz/ormx"
	"sync"
	"video_web/internal/model/entity"
)

type noticeRepo struct {
	ormx.IRepo[entity.Notice]
}

var noticeOnce sync.Once
var _notice *noticeRepo

func Notice() *noticeRepo {
	noticeOnce.Do(func() {
		_notice = &noticeRepo{IRepo: ormx.NewRepo[entity.Notice]()}
	})
	return _notice
}
