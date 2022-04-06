package logic

import (
	"context"
	"video_web/internal/model"
	"video_web/internal/repo"
)

type CategoryLogic struct {
	categoryRepo *repo.CategoryRepo
}

func NewCategoryLogic(categoryRepo *repo.CategoryRepo) *CategoryLogic {
	return &CategoryLogic{categoryRepo: categoryRepo}
}

func (c CategoryLogic) Add(ctx context.Context, category *model.Category) error {
	return c.categoryRepo.Add(ctx, category)
}

func (c CategoryLogic) List(ctx context.Context) ([]*model.Category, error) {
	return c.categoryRepo.GetList(ctx)
}
