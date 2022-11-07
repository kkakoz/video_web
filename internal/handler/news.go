package handler

import (
	"github.com/labstack/echo"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
)

type newsfeedHandler struct {
}

func (newsfeedHandler) Add(ctx echo.Context) error {
	req := &dto.NewsfeedAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Newsfeed().Add(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (newsfeedHandler) UserNews(ctx echo.Context) error {
	req := &dto.UserId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	data, err := logic.Newsfeed().UserNews(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"data": data,
	})
}

func (newsfeedHandler) List(ctx echo.Context) error {
	req := &dto.NewsfeedList{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := logic.Newsfeed().List(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"data": list,
	})
}
