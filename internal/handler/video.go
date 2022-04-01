package handler

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"video_web/internal/domain"
	"video_web/internal/dto/request"
	"video_web/internal/pkg/mdctx"
	"video_web/pkg/errno"
)

type VideoHandler struct {
	videoLogic domain.IVideoLogic
}

func NewVideoHandler(videoLogic domain.IVideoLogic) *VideoHandler {
	return &VideoHandler{videoLogic: videoLogic}
}

func (v VideoHandler) AddVideo(ctx echo.Context) error {
	req := &request.VideoAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = v.videoLogic.Add(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (v VideoHandler) AddVideoEpisode(ctx echo.Context) error {
	req := &request.EpisodeAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	episode := &domain.Episode{}
	err = copier.Copy(episode, req)
	if err != nil {
		return errno.New400("参数错误")
	}
	err = v.videoLogic.AddEpisode(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (v VideoHandler) AddEpisode(ctx echo.Context) error {
	req := &request.EpisodeAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	episode := &domain.Episode{}
	err = copier.Copy(episode, req)
	if err != nil {
		return errno.New400("参数错误")
	}
	err = v.videoLogic.AddEpisode(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (v VideoHandler) GetVideo(ctx echo.Context) error {
	req := &request.VideoIdReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	video, err := v.videoLogic.GetVideo(mdctx.NewCtx(ctx.Request()), req.VideoId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, video)
}

func (v VideoHandler) GetVideos(ctx echo.Context) error {
	req := &request.VideosReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	videos, err := v.videoLogic.GetVideos(mdctx.NewCtx(ctx.Request()), req.CategoryId, req.LastValue, req.OrderType)
	if err != nil {
		return err
	}
	return ctx.JSON(200, videos)
}

func (v VideoHandler) DelVideo(ctx echo.Context) error {
	req := &request.EpisodeIdReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = v.videoLogic.DelVideoEpisode(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}
