package router

import (
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/handler"
)

type userRouter struct {
	handler *handler.UserHandler
}

func NewUserRouter(handler *handler.UserHandler) *userRouter {
	return &userRouter{handler: handler}
}

func (u userRouter) AddRouter(e *echo.Echo) {
	userG := e.Group("/users")
	{
		userG.POST("", u.handler.Register)
		userG.POST("/login", u.handler.Login)
		userG.GET("/cur", u.handler.GetCurUser)
	}

	{
		userG.GET("/:user_id", u.handler.GetUser, authority)
	}
}
