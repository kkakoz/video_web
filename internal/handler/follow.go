package handler

import (
	"github.com/labstack/echo"
	"video_web/internal/dto/request"
	"video_web/internal/logic"
	"video_web/pkg/local"
)

type FollowHandler struct {
	followLogic *logic.FollowLogic
}

func NewFollowHandler(followLogic *logic.FollowLogic) *FollowHandler {
	return &FollowHandler{followLogic: followLogic}
}

func (item *FollowHandler) Follow(ctx echo.Context) error {
	req := &request.FollowReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = item.followLogic.Follow(local.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return nil
}

func (item *FollowHandler) Fans(ctx echo.Context) error {
	req := &request.FollowFansReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := item.followLogic.Fans(local.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *FollowHandler) Followers(ctx echo.Context) error {
	req := &request.FollowersReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := item.followLogic.Followers(local.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *FollowHandler) IsFollowed(ctx echo.Context) error {
	req := &request.FollowIsReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := item.followLogic.IsFollow(local.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *FollowHandler) GroupAdd(ctx echo.Context) error {
	req := &request.FollowGroupAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = item.followLogic.GroupAdd(local.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return nil
}

func (item *FollowHandler) Groups(ctx echo.Context) error {
	list, err := item.followLogic.Groups(local.NewCtx(ctx.Request()))
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}
