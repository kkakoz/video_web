package handler

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/dto/request"
	"video_web/internal/logic"
	"video_web/internal/model"
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

func (item *videoHandler) AddVideo(ctx echo.Context) error {
	req := &request.VideoAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Video().Add(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *videoHandler) AddVideoEpisode(ctx echo.Context) error {
	req := &request.EpisodeAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	episode := &model.Episode{}
	err = copier.Copy(episode, req)
	if err != nil {
		return errno.New400("参数错误")
	}
	err = logic.Video().AddEpisode(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *videoHandler) AddEpisode(ctx echo.Context) error {
	req := &request.EpisodeAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	episode := &model.Episode{}
	err = copier.Copy(episode, req)
	if err != nil {
		return errno.New400("参数错误")
	}
	err = logic.Video().AddEpisode(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *videoHandler) Get(ctx echo.Context) error {
	req := &request.VideoIdReq{}
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
	req := &request.VideosReq{}
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
	req := &request.BackVideosReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	videos, _, err := logic.Video().GetBackList(ctx.Request().Context(), req.CategoryId, req.OrderType, req.Pager)
	if err != nil {
		return err
	}
	return ctx.JSON(200, videos)
}

func (item *videoHandler) DelVideo(ctx echo.Context) error {
	req := &request.EpisodeIdReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Video().DelVideoEpisode(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *videoHandler) Ws(ctx echo.Context) error {
	req := &request.VideoIdReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	return ws.VideoConn().Add(ctx.Response(), ctx.Request(), req.VideoId)
}
