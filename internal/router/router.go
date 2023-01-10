package router

import (
	"github.com/kkakoz/pkg/echox"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func NewHttp() http.Handler {
	e := echo.New()
	e.Binder = echox.NewBinder()
	e.Validator = echox.NewValidator()
	e.HTTPErrorHandler = ErrHandler()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	// e.Use(setAccessOriginUrl)

	e.Debug = true
	AppRouter(e)
	BackRouter(e)
	return e
}
