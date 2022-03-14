package domain

import (
	"context"

	"github.com/kkakoz/ormx"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ICategoryLogic interface {
	Add(ctx context.Context, category *Category) error
	List(ctx context.Context) ([]*Category, error)
}

type ICategoryRepo interface {
	ormx.IRepo[Category]
}
