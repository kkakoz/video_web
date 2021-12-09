package repo

import (
	"context"
	"github.com/pkg/errors"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/pkg/mysqlx"
)

var _ domain.ICategoryRepo = (*CategoryRepo)(nil)

type CategoryRepo struct {

}

func NewCategoryRepo() domain.ICategoryRepo {
	return &CategoryRepo{}
}

func (c CategoryRepo) Add(ctx context.Context, category *domain.Category) error {
	db := mysqlx.GetDB(ctx)
	err := db.Create(category).Error
	return errors.Wrap(err, "添加失败")
}

func (c CategoryRepo) List(ctx context.Context) ([]*domain.Category, error) {
	db := mysqlx.GetDB(ctx)
	list := make([]*domain.Category, 0)
	err := db.Find(&list).Error
	return list, err
}

