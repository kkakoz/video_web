package domain

import "context"

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ICategoryLogic interface {
	Add(ctx context.Context, category *Category) error
	List(ctx context.Context) ([]*Category, error)
}

type ICategoryRepo interface {
	Add(ctx context.Context, category *Category) error
	List(ctx context.Context) ([]*Category, error)
}