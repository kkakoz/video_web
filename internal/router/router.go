package router

import (
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/handler"
	"github/kkakoz/video_web/pkg/echox"
)

func NewRouter(user handler.UserHandler) *echo.Echo {
	e := echo.New()
	e.Binder = echox.NewBinder()
	e.Validator = echox.NewValidator()
	userRouter(e, user)
	return e
}
