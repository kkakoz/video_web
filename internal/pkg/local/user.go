package local

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"video_web/internal/model"
	"video_web/pkg/errno"
	"video_web/pkg/safe"
)

const UserLocalKey = "user:local:key"

func GetUser(ctx context.Context) (*model.User, error) {
	md, b := metadata.FromIncomingContext(ctx)
	if !b {
		return nil, errors.New("user not found")
	}
	user := &model.User{}
	err := json.Unmarshal([]byte(safe.Get(func() string {
		return md.Get(UserLocalKey)[0]
	})), user)
	if err != nil || user.ID == 0 {
		return nil, errno.NewErr(401, 401, "user not found")
	}
	return user, nil
}

type userLocalKey struct{}

func WithUserLocal(ctx context.Context, value *model.User) context.Context {
	return context.WithValue(ctx, userLocalKey{}, value)
}

func GetUserLocal(ctx context.Context) (*model.User, bool) {
	v, ok := ctx.Value(userLocalKey{}).(*model.User)
	return v, ok
}
