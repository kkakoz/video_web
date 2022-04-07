package handler

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"video_web/internal/dto/request"
	"video_web/internal/logic"
	"video_web/internal/model"
	"video_web/internal/pkg/mdctx"
	"video_web/internal/pkg/ws"
	"video_web/pkg/errno"
)

type VideoHandler struct {
	videoLogic *logic.VideoLogic
	videoConn  *ws.VideoConn
}

func NewVideoHandler(videoLogic *logic.VideoLogic, videoConn *ws.VideoConn) *VideoHandler {
	return &VideoHandler{videoLogic: videoLogic, videoConn: videoConn}
}

func (item *VideoHandler) AddVideo(ctx echo.Context) error {
	req := &request.VideoAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = item.videoLogic.Add(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *VideoHandler) AddVideoEpisode(ctx echo.Context) error {
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
	err = item.videoLogic.AddEpisode(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *VideoHandler) AddEpisode(ctx echo.Context) error {
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
	err = item.videoLogic.AddEpisode(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *VideoHandler) GetVideo(ctx echo.Context) error {
	req := &request.VideoIdReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	video, err := item.videoLogic.GetVideo(mdctx.NewCtx(ctx.Request()), req.VideoId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, video)
}

func (item *VideoHandler) GetVideos(ctx echo.Context) error {
	req := &request.VideosReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	videos, err := item.videoLogic.GetVideos(mdctx.NewCtx(ctx.Request()), req.CategoryId, req.LastValue, req.OrderType)
	if err != nil {
		return err
	}
	return ctx.JSON(200, videos)
}

func (item *VideoHandler) DelVideo(ctx echo.Context) error {
	req := &request.EpisodeIdReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = item.videoLogic.DelVideoEpisode(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *VideoHandler) Ws(ctx echo.Context) error {
	req := &request.VideoIdReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	return item.videoConn.Add(ctx.Response(), ctx.Request(), req.VideoId)
}
