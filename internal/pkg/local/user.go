package local

import (
	"context"
	"video_web/internal/model/entity"
	"video_web/pkg/errno"
)

const UserLocalKey = "user:local:key"

func GetUser(ctx context.Context) (*entity.User, error) {
	v, ok := ctx.Value(userLocalKey{}).(*entity.User)
	if !ok {
		return nil, errno.NewErr(401, 401, "用户未登录")
	}
	return v, nil
	//md, b := metadata.FromIncomingContext(ctx)
	//if !b {
	//	return nil, errno.NewErr(401, 401, "用户未登录")
	//}
	//user := &entity.User{}
	//err := json.Unmarshal([]byte(safe.Get(func() string {
	//	return md.Get(UserLocalKey)[0]
	//})), user)
	//if err != nil || user.ID == 0 {
	//	return nil, errno.NewErr(401, 401, "用户未登录")
	//}
	//return user, nil
}

type userLocalKey struct{}

func WithUserLocal(ctx context.Context, value *entity.User) context.Context {
	return context.WithValue(ctx, userLocalKey{}, value)
}

func GetUserLocal(ctx context.Context) (*entity.User, bool) {
	v, ok := ctx.Value(userLocalKey{}).(*entity.User)
	return v, ok
}
