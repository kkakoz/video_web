package logic

import (
	"context"
	"sync"
	"video_web/internal/logic/internal/repo"
	"video_web/internal/model/entity"
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

func (categoryLogic) Add(ctx context.Context, category *entity.Category) error {
	return repo.Category().Add(ctx, category)
}

func (categoryLogic) List(ctx context.Context) ([]*entity.Category, error) {
	return repo.Category().GetList(ctx)
}
