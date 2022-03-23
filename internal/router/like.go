package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

type likeRouter struct {
	handler *handler.LikeHandler
}

func NewLikeRouter(handler *handler.LikeHandler) *likeRouter {
	return &likeRouter{handler: handler}
}

func (item likeRouter) AddRouter(e *echo.Echo) {
	likeG := e.Group("/api/likes")
	{
		likeG.POST("", item.handler.Like, authority)
		likeG.GET("", item.handler.IsLike, authority)
	}
}
