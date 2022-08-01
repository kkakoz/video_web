package handler

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/ws"
	"video_web/pkg/errno"
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
	episode := &entity.Video{}
	err = copier.Copy(episode, req)
	if err != nil {
		return errno.New400("参数错误")
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
	return ctx.JSON(200, videos)
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
	return ctx.JSON(200, nil)
}

func (item *videoHandler) Ws(ctx echo.Context) error {
	req := &dto.VideoId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	return ws.VideoConn().Add(ctx.Response(), ctx.Request(), req.VideoId)
}
