package router

import (
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/handler"
	"github/kkakoz/video_web/pkg/echox"
	"go.uber.org/fx"
	"net/http"
)

func NewHttp(user *handler.UserHandler) http.Handler {
	e := echo.New()
	e.Binder = echox.NewBinder()
	e.Validator = echox.NewValidator()
	userRouter(e, user)
	return e
}

var Provider = fx.Provide(NewHttp)