package logic

import (
	"context"
	"video_web/internal/domain"
)

var _ domain.ICategoryLogic = (*CategoryLogic)(nil)

type CategoryLogic struct {
	categoryRepo domain.ICategoryRepo
}

func NewCategoryLogic(categoryRepo domain.ICategoryRepo) domain.ICategoryLogic {
	return &CategoryLogic{categoryRepo: categoryRepo}
}

func (c CategoryLogic) Add(ctx context.Context, category *domain.Category) error {
	return c.categoryRepo.Add(ctx, category)
}

func (c CategoryLogic) List(ctx context.Context) ([]*domain.Category, error) {
	return c.categoryRepo.List(ctx)
}
