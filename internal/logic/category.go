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

func (item *CategoryLogic) Add(ctx context.Context, category *model.Category) error {
	return item.categoryRepo.Add(ctx, category)
}

func (item *CategoryLogic) List(ctx context.Context) ([]*model.Category, error) {
	return item.categoryRepo.GetList(ctx)
}
