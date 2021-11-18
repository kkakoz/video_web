package handler

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/internal/dto"
	"google.golang.org/grpc/metadata"
)

type VideoHandler struct {
	videoLogic domain.IVideoLogic
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
	copier.Copy(video, req)
	err = v.videoLogic.AddVideo(newCtx, video)
	if err != nil {
		return err
	}
	return nil
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
		return err
	}
	err = v.videoLogic.AddEpisode(newCtx, episode)
	if err != nil {
		return err
	}
	return nil
}

