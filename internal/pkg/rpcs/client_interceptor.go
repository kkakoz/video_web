package rpcs

import (
	"context"
	"encoding/json"
	userpb "github.com/kkakoz/video-rpc/pb/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"video_web/internal/pkg/local"
)

func userInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if value, exist := local.GetUserExist(ctx); exist {
			u := &userpb.UserInfoRes{
				ID:          value.ID,
				Type:        uint32(value.Type),
				Name:        value.Name,
				Avatar:      value.Avatar,
				Brief:       value.Brief,
				FollowCount: value.FollowCount,
				FansCount:   value.FansCount,
				LikeCount:   value.LikeCount,
				State:       value.State,
				LastLogin:   timestamppb.New(value.LastLogin.Time),
				Email:       value.Email,
				Phone:       value.Phone,
				CreatedAt:   timestamppb.New(value.CreatedAt.Time),
				UpdatedAt:   timestamppb.New(value.UpdatedAt.Time),
				DeletedAt:   timestamppb.New(value.DeletedAt.Time),
			}
			userBytes, _ := json.Marshal(u)
			ctx = metadata.NewIncomingContext(ctx, metadata.New(map[string]string{
				"user": string(userBytes),
			}))
		}
		err := invoker(ctx, method, req, reply, cc)
		return err
	}
}
