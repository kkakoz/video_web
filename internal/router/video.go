package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

type videoRouter struct {
	handler *handler.VideoHandler
}

func NewVideoRouter(handler *handler.VideoHandler) *videoRouter {
	return &videoRouter{handler: handler}
}

func (v videoRouter) AddRouter(e *echo.Echo) {
	videoG := e.Group("/api/videos")
	{
		videoG.POST("", v.handler.AddVideo, authority)
		videoG.POST("/:video_id/episodes", v.handler.AddVideoEpisode, authority)
		videoG.GET("/:video_id", v.handler.GetVideo)
		videoG.GET("", v.handler.GetVideos)
	}

	episodeG := e.Group("/api/episodes", authority)
	{
		episodeG.POST("", v.handler.AddEpisode)
		episodeG.DELETE("/:episode_id", v.handler.DelVideo)
	}

}
