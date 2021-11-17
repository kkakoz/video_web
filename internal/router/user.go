package router

import (
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/handler"
)

func userRouter(e *echo.Echo, handler handler.UserHandler) {
	user := e.Group("/user")

	user.POST("", handler.Register)
	user.GET("", handler.GetUser)
	user.POST("/login", handler.Login)
}
