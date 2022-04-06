package repo

import (
	"video_web/internal/model"

	"github.com/kkakoz/ormx"
)

type CategoryRepo struct {
	ormx.IRepo[model.Category]
}

func NewCategoryRepo() *CategoryRepo {
	return &CategoryRepo{
		ormx.NewRepo[model.Category](),
	}
}
