package handler

import (
	"github.com/kkakoz/pkg/echox"
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
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
	req := &dto.VideoAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	video, err := logic.Video().Add(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, video)
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
	video, err := logic.Video().GetById(ctx.Request().Context(), req.VideoId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, video)
}

func (item *videoHandler) AddView(ctx echo.Context) error {
	req := &dto.VideoId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Video().AddView(ctx.Request().Context(), req.VideoId)
	if err != nil {
		return err
	}
	return echox.OK(ctx)
}

func (item *videoHandler) List(ctx echo.Context) error {
	req := &dto.Videos{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	videos, err := logic.Video().List(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]interface{}{
		"data": videos,
	})
}

func (item *videoHandler) UserVideoState(ctx echo.Context) error {
	req := &dto.VideoId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	userState, err := logic.Video().UserState(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, userState)
}

func (item *videoHandler) Recommend(ctx echo.Context) error {
	req := &dto.VideoId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	userState, err := logic.Video().Recommend(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, userState)
}

func (item *videoHandler) Rankings(ctx echo.Context) error {
	req := &dto.Rankings{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	data, err := logic.Video().Rankings(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"data": data,
	})
}
