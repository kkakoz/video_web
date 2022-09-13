package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/internal/pkg/ws"
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
	req := &dto.VideoAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Video().AddVideo(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *videoHandler) Get(ctx echo.Context) error {
	req := &dto.VideoId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	video, err := logic.Video().GetVideo(ctx.Request().Context(), req.VideoId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, video)
}

func (item *videoHandler) GetList(ctx echo.Context) error {
	req := &dto.Videos{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	videos, err := logic.Video().GetVideos(ctx.Request().Context(), req.CategoryId, req.LastValue, req.OrderType)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"data": videos,
	})
}

func (item *videoHandler) GetBackList(ctx echo.Context) error {
	req := &dto.BackVideoList{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	videos, count, err := logic.Video().GetPageList(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"count": count,
		"items": videos,
	})
}

func (item *videoHandler) Del(ctx echo.Context) error {
	req := &dto.VideoId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Video().DelVideo(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return echox.OK(ctx)
}

func (item *videoHandler) Ws(ctx echo.Context) error {
	req := &dto.VideoId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	return ws.VideoConn().Add(ctx.Response(), ctx.Request(), req.VideoId)
}

func (item *videoHandler) UserVideoState(ctx echo.Context) error {
	req := &dto.CollectionId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	userState, err := logic.Collection().UserState(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, userState)
}
