package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

type ossRouter struct {
	handler *handler.OssHandler
}

func NewOssRouter(handler *handler.OssHandler) *ossRouter {
	return &ossRouter{handler: handler}
}

func (o ossRouter) AddRouter(e *echo.Echo) {
	OssG := e.Group("/api/oss")

	OssG.GET("/conf", o.handler.GetConf)
}
