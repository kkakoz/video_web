package logic

import (
	"context"
	"sync"
	"video_web/internal/model"
	"video_web/internal/repo"
)

type categoryLogic struct{}

var categoryOnce sync.Once
var _category *categoryLogic

func Category() *categoryLogic {
	categoryOnce.Do(func() {
		_category = &categoryLogic{}
	})
	return _category
}

func (item *categoryLogic) Add(ctx context.Context, category *model.Category) error {
	return repo.Category().Add(ctx, category)
}

func (item *categoryLogic) List(ctx context.Context) ([]*model.Category, error) {
	return repo.Category().GetList(ctx)
}
