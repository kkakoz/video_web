package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/dto/request"
	"video_web/internal/logic"
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
	req := &request.LikeReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Like().Like(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return nil
}

func (item *likeHandler) IsLike(ctx echo.Context) error {
	req := &request.LikeIsReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	b, err := logic.Like().IsLike(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, b)
}
