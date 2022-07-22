package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/dto/request"
	"video_web/internal/logic"
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
	req := &request.FollowReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Follow().Follow(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return nil
}

func (item *followHandler) Fans(ctx echo.Context) error {
	req := &request.FollowFansReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := logic.Follow().Fans(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *followHandler) Followers(ctx echo.Context) error {
	req := &request.FollowersReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := logic.Follow().Followers(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *followHandler) IsFollowed(ctx echo.Context) error {
	req := &request.FollowIsReq{}
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
	req := &request.FollowGroupAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Follow().GroupAdd(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return nil
}

func (item *followHandler) Groups(ctx echo.Context) error {
	list, err := logic.Follow().Groups(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}
