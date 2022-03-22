package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

type categoryRouter struct {
	handler *handler.CategoryHandler
}

func NewCategoryRouter(handler *handler.CategoryHandler) *categoryRouter {
	return &categoryRouter{handler: handler}
}

func (c categoryRouter) AddRouter(e *echo.Echo) {
	categoryG := e.Group("/api/categories")
	{
		categoryG.POST("", c.handler.Add, authority)
		categoryG.GET("", c.handler.List)
	}
}
