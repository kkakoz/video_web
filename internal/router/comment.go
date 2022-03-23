package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

type commentRouter struct {
	handler *handler.CommentHandler
}

func NewCommentRouter(handler *handler.CommentHandler) *commentRouter {
	return &commentRouter{handler: handler}
}

func (item commentRouter) AddRouter(e *echo.Echo) {
	commentG := e.Group("/api/comments")
	{
		commentG.POST("", item.handler.Add, authority)
		commentG.POST("/:comment_id/sub_comments", item.handler.AddSubComment, authority)
		commentG.GET("", item.handler.Get)
		commentG.GET("/:comment_id/sub_comments", item.handler.GetSubComment)
	}
}
