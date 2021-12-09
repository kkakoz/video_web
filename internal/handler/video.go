package handler

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/internal/dto"
	"github/kkakoz/video_web/pkg/errno"
	"google.golang.org/grpc/metadata"
)

type VideoHandler struct {
	videoLogic domain.IVideoLogic
}

func NewVideoHandler(videoLogic domain.IVideoLogic) *VideoHandler {
	return &VideoHandler{videoLogic: videoLogic}
}

func (v VideoHandler) AddVideo(ctx echo.Context) error {
	md := metadata.New(nil)
	newCtx := metadata.NewIncomingContext(ctx.Request().Context(), md)
	req := &dto.AddVideoReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	video := &domain.Video{}
	err = copier.Copy(video, req)
	if err != nil {
		return errno.New400("参数错误")
	}
	err = v.videoLogic.AddVideo(newCtx, video)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (v VideoHandler) AddEpisode(ctx echo.Context) error {
	md := metadata.New(nil)
	newCtx := metadata.NewIncomingContext(ctx.Request().Context(), md)
	req := &dto.AddEpisodeReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	episode := &domain.Episode{}
	err = copier.Copy(episode, req)
	if err != nil {
		return errno.New400("参数错误")
	}
	err = v.videoLogic.AddEpisode(newCtx, episode)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (v VideoHandler) GetVideo(ctx echo.Context) error {
	md := metadata.New(nil)
	newCtx := metadata.NewIncomingContext(ctx.Request().Context(), md)
	req := &dto.VideoIdReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	video, err := v.videoLogic.GetVideo(newCtx, req.VideoId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, video)
}

func (v VideoHandler) DelVideo(ctx echo.Context) error {
	md := metadata.New(nil)
	newCtx := metadata.NewIncomingContext(ctx.Request().Context(), md)
	req := &dto.EpisodeIdReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = v.videoLogic.DelEpisode(newCtx, req.EpisodeId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}
