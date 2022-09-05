package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/pkg/echox"
)

type videoHandler struct {
}

var videoOnce sync.Once
var _video *videoHandler

func Video() *videoHandler {
	videoOnce.Do(func() {
		_video = &videoHandler{}
	})
	return _video
}

func (item *videoHandler) Add(ctx echo.Context) error {
	req := &dto.CollectionAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Video().Add(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return echox.OK(ctx)
}

func (item *videoHandler) BackList(ctx echo.Context) error {
	req := &dto.BackCollectionList{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	collections, count, err := logic.Video().GetPageList(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"count": count,
		"data":  collections,
	})
}

func (item *videoHandler) Get(ctx echo.Context) error {
	req := &dto.VideoId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	resources, err := logic.Video().Get(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, resources)
}

func (item *videoHandler) List(ctx echo.Context) error {
	req := &dto.Videos{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	resources, err := logic.Video().List(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, resources)
}
