package repo

import (
	"video_web/internal/domain"

	"github.com/kkakoz/ormx"
)

var _ domain.ICategoryRepo = (*CategoryRepo)(nil)

type CategoryRepo struct {
	ormx.Repo[domain.Category]
}

func NewCategoryRepo() domain.ICategoryRepo {
	return &CategoryRepo{}
}
