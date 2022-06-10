package repo

import (
	"github.com/kkakoz/ormx"
	"video_web/internal/model"
)

type DanmuRepo struct {
	ormx.IRepo[model.Danmu]
}

func NewDanmuRepo() *DanmuRepo {
	return &DanmuRepo{
		ormx.NewRepo[model.Danmu](),
	}
}
