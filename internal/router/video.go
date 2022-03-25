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
		videoG.POST("/:video_id/episodes", v.handler.AddEpisode, authority)
		videoG.GET("/:video_id", v.handler.GetVideo)
		videoG.GET("", v.handler.GetVideos)
	}

	episodeG := e.Group("/episodes", authority)
	{
		episodeG.DELETE("/:episode_id", v.handler.DelVideo)
	}

}
