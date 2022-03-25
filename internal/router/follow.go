package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

type followRouter struct {
	handler *handler.FollowHandler
}

func NewFollowRouter(handler *handler.FollowHandler) *followRouter {
	return &followRouter{handler: handler}
}

func (item followRouter) AddRouter(e *echo.Echo) {
	followG := e.Group("/api/follow")
	{
		followG.POST("", item.handler.Follow, authority)
		followG.GET("/fans", item.handler.Fans)
		followG.GET("/followers", item.handler.Followers, authority)
		followG.GET("/followed", item.handler.IsFollowed, authority)
		followG.POST("/groups", item.handler.GroupAdd, authority)
		followG.GET("/groups", item.handler.Groups, authority)
	}
}
