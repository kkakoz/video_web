package async

import (
	"context"
	"github.com/kkakoz/ormx"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
)

type handler func(event *dto.Event) error

var handlerMap = map[dto.EventType]handler{}

func init() {
	handlerMap[dto.EventTypeLike] = handlerLike
	handlerMap[dto.EventTypeUserRegister] = handlerUserRegister
}

func handlerLike(event *dto.Event) error {
	return nil
}

func handlerUserRegister(event *dto.Event) error {
	return ormx.Transaction(context.Background(), func(txctx context.Context) error {
		return logic.User().UserInit(txctx, event.ActorId)
	})
}
