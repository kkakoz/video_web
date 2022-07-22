package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

func BackRouter(e *echo.Echo) {
	back := e.Group("/api/back", authority)

	{
		videoG := back.Group("/videos")

		videoG.GET("/", handler.Video().GetBackList)
		videoG.GET("/:video_id", handler.Video().Get)

		userG := back.Group("/users")
		{
			userG.GET("/local", handler.User().GetCurUser)
		}

		commentG := e.Group("/comments")
		{
			commentG.GET("", handler.Comment().Get)
			commentG.GET("/:comment_id/sub_comments", handler.Comment().GetSubComment)
			commentG.DELETE("/:comment_id", handler.Comment().Delete)
		}

		subCommentG := e.Group("/sub/comments")
		{
			subCommentG.DELETE("/sub_comments/:sub_comment_id", handler.Comment().DeleteSubComment)
		}
	}
}
