package repo

import (
	"sync"
	"video_web/internal/model/entity"

	"github.com/kkakoz/ormx"
)

type categoryRepo struct {
	ormx.IRepo[entity.Category]
}

var categoryOnce sync.Once
var _category *categoryRepo

func Category() *categoryRepo {
	categoryOnce.Do(func() {
		_category = &categoryRepo{IRepo: ormx.NewRepo[entity.Category]()}
	})
	return _category
}
