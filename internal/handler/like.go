package handler

import (
	"github.com/labstack/echo"
	"video_web/internal/domain"
	"video_web/internal/dto/request"
	"video_web/pkg/local"
)

type LikeHandler struct {
	likeLogic domain.ILikeLogic
}

func NewLikeHandler(likeLogic domain.ILikeLogic) *LikeHandler {
	return &LikeHandler{likeLogic: likeLogic}
}

func (item *LikeHandler) Like(ctx echo.Context) error {
	req := &request.LikeReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = item.likeLogic.Like(local.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return nil
}

func (item *LikeHandler) IsLike(ctx echo.Context) error {
	req := &request.LikeIsReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	b, err := item.likeLogic.IsLike(local.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, b)
}
