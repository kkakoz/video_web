package router

import (
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/handler"
)

type videoRouter struct {
	handler *handler.VideoHandler
}

func NewVideoRouter(handler *handler.VideoHandler) *videoRouter {
	return &videoRouter{handler: handler}
}

func (v videoRouter) AddRouter(e *echo.Echo) {
	videoG := e.Group("/videos", authority)
	{
		videoG.POST("", v.handler.AddVideo)
		videoG.POST("/:video_id/episodes", v.handler.AddEpisode)
		videoG.GET("/:video_id", v.handler.GetVideo)
	}

	episodeG := e.Group("/episodes", authority)
	{
		episodeG.DELETE("/:episode_id", v.handler.DelVideo)
	}

}
