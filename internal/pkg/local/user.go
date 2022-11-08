package local

import (
	"context"
	"video_web/internal/model/entity"
	"video_web/pkg/errno"
)

const UserLocalKey = "user:local:key"

type local[T any] struct {
	key string
}

func newLocal[T any](key string) *local[T] {
	return &local[T]{key: key}
}

func (local[T]) Get(ctx context.Context) (*T, error) {
	v, ok := ctx.Value(userLocalKey{}).(*T)
	if !ok {
		return nil, errno.NewErr(401, 401, "not found")
	}
	return v, nil
}

func (local[T]) GetExist(ctx context.Context) bool {
	_, ok := ctx.Value(userLocalKey{}).(*T)
	return ok
}

func GetUser(ctx context.Context) (*entity.User, error) {
	v, ok := ctx.Value(userLocalKey{}).(*entity.User)
	if !ok {
		return nil, errno.NewErr(401, 401, "用户未登录")
	}
	return v, nil
}

type userLocalKey struct{}

func WithUserLocal(ctx context.Context, value *entity.User) context.Context {
	return context.WithValue(ctx, userLocalKey{}, value)
}
