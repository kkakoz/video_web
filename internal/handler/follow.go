package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/pkg/echox"
)

type followHandler struct {
}

var followOnce sync.Once
var _follow *followHandler

func Follow() *followHandler {
	followOnce.Do(func() {
		_follow = &followHandler{}
	})
	return _follow
}

func (item *followHandler) Follow(ctx echo.Context) error {
	req := &dto.Follow{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Follow().Follow(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return echox.OK(ctx)
}

func (item *followHandler) Fans(ctx echo.Context) error {
	req := &dto.FollowFans{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := logic.Follow().Fans(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"data": list,
	})
}

func (item *followHandler) Followers(ctx echo.Context) error {
	req := &dto.Followers{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := logic.Follow().Followers(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"data": list,
	})
}

func (item *followHandler) IsFollowed(ctx echo.Context) error {
	req := &dto.FollowIs{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := logic.Follow().IsFollow(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *followHandler) GroupAdd(ctx echo.Context) error {
	req := &dto.FollowGroupAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Follow().GroupAdd(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return echox.OK(ctx)
}

func (item *followHandler) Groups(ctx echo.Context) error {
	list, err := logic.Follow().Groups(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}
