package repo

import (
	"sync"
	"video_web/internal/model"

	"github.com/kkakoz/ormx"
)

type categoryRepo struct {
	ormx.IRepo[model.Category]
}

var categoryOnce sync.Once
var _category *categoryRepo

func Category() *categoryRepo {
	categoryOnce.Do(func() {
		_category = &categoryRepo{IRepo: ormx.NewRepo[model.Category]()}
	})
	return _category
}
