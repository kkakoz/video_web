package repo

import (
	"sync"
	"video_web/internal/model/entity"

	"github.com/kkakoz/ormx"
)

type resourceRepo struct {
	ormx.IRepo[entity.Resource]
}

var resourceOnce sync.Once
var _resource *resourceRepo

func Resource() *resourceRepo {
	resourceOnce.Do(func() {
		_resource = &resourceRepo{IRepo: ormx.NewRepo[entity.Resource]()}
	})
	return _resource
}
