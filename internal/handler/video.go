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
	req := &request.AddVideoReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	video := &domain.Video{}
	err = copier.Copy(video, req)
	if err != nil {
		return errno.New400("参数错误")
	}
	err = v.videoLogic.AddVideo(mdctx.NewCtx(ctx.Request()), video)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (v VideoHandler) AddEpisode(ctx echo.Context) error {
	req := &request.AddEpisodeReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	episode := &domain.Episode{}
	err = copier.Copy(episode, req)
	if err != nil {
		return errno.New400("参数错误")
	}
	err = v.videoLogic.AddEpisode(mdctx.NewCtx(ctx.Request()), episode)
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

func (v VideoHandler) DelVideo(ctx echo.Context) error {
	req := &request.EpisodeIdReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = v.videoLogic.DelEpisode(mdctx.NewCtx(ctx.Request()), req.EpisodeId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}