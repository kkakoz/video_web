package handler

import (
	"github.com/kkakoz/pkg/echox"
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
)

type newsfeedHandler struct {
}

var newsfeedOnce sync.Once
var _newsfeed *newsfeedHandler

func Newsfeed() *newsfeedHandler {
	newsfeedOnce.Do(func() {
		_newsfeed = &newsfeedHandler{}
	})
	return _newsfeed
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
	return echox.Ok(ctx)
}

func (newsfeedHandler) UserNews(ctx echo.Context) error {
	req := &dto.UserNewsfeedList{}
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

func (newsfeedHandler) FollowedNewsFeeds(ctx echo.Context) error {
	req := &dto.LastId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	data, err := logic.Newsfeed().FollowedNewsFeeds(ctx.Request().Context(), req)
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
