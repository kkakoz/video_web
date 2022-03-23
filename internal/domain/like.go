package domain

import (
	"context"
	"github.com/kkakoz/ormx"
	"video_web/internal/dto/request"
)

type Like struct {
	UserId     int64 `json:"user_id"`
	TargetType uint8 `json:"target_type"`
	TargetId   int64 `json:"target_id"`
}

type ILikeLogic interface {
	Like(ctx context.Context, req *request.LikeReq) error
	IsLike(ctx context.Context, req *request.LikeIsReq) (bool, error)
}

type ILikeRepo interface {
	ormx.IRepo[Like]
}
