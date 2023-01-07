package handler

import (
	"github.com/kkakoz/pkg/echox"
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
)

type likeHandler struct {
}

var likeOnce sync.Once
var _like *likeHandler

func Like() *likeHandler {
	likeOnce.Do(func() {
		_like = &likeHandler{}
	})
	return _like
}

func (item *likeHandler) Like(ctx echo.Context) error {
	req := &dto.Like{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Like().Like(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return echox.Ok(ctx)
}
