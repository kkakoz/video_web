package router

import (
	"video_web/internal/handler"

	"github.com/labstack/echo"
)

type userRouter struct {
	handler *handler.UserHandler
}

func NewUserRouter(handler *handler.UserHandler) *userRouter {
	return &userRouter{handler: handler}
}

func (u userRouter) AddRouter(e *echo.Echo) {
	userG := e.Group("/api/users")
	{
		userG.POST("", u.handler.Register)
		userG.POST("/login", u.handler.Login)
		userG.GET("/local", u.handler.GetCurUser)
	}

	{
		userG.GET("/:user_id", u.handler.GetUser, authority)
	}
}
