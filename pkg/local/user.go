package local

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"video_web/pkg/errno"
	"video_web/pkg/safe"
)

const UserLocalKey = "user:local:key"

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Brief       string `json:"brief"`
	FollowCount int64  `json:"follow_count"`
	FansCount   int64  `json:"fans_count"`
	LikeCount   int64  `json:"like_count"`
	State       int32  `json:"state"`
	LastLogin   int64  `json:"last_login"`
}

func GetUser(ctx context.Context) (*User, error) {
	md, b := metadata.FromIncomingContext(ctx)
	if !b {
		return nil, errors.New("user not found")
	}
	user := &User{}
	err := json.Unmarshal([]byte(safe.Get(func() string {
		return md.Get(UserLocalKey)[0]
	})), user)
	if err != nil || user.ID == 0 {
		return nil, errno.NewErr(401, 401, "user not found")
	}
	return user, nil
}
