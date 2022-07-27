package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

func BackRouter(e *echo.Echo) {
	back := e.Group("/api/back")

	{
		userG := back.Group("/users")
		userG.POST("/login", handler.User().Login)
	}

	authBack := e.Group("/api/back", authority)

	{
		categoryG := authBack.Group("/categories")
		{
			categoryG.GET("", handler.Category().List)
		}

		collectionG := authBack.Group("/collections")
		{
			// 添加合集
			collectionG.POST("", handler.Video().AddCollection)
		}

		videoG := authBack.Group("/videos")
		{
			// 添加视频
			videoG.POST("", handler.Video().AddVideo)
			// 删除视频
			videoG.DELETE("/:video_id", handler.Video().DelVideo)
			// 视频列表
			videoG.GET("", handler.Video().GetBackList)
		}

		userG := authBack.Group("/users")
		{
			// 当前用户
			userG.GET("/local", handler.User().GetCurUser)
		}

		OssG := authBack.Group("/oss")

		OssG.GET("/conf", handler.Oss().GetConf)

		commentG := e.Group("/comments")
		{
			// 获取评论
			commentG.GET("", handler.Comment().Get)
			// 子评论
			commentG.GET("/:comment_id/sub_comments", handler.Comment().GetSubComment)
			// 删除评论
			commentG.DELETE("/:comment_id", handler.Comment().Delete)
			commentG.DELETE("/:comment_id/sub_comments/:sub_comment_id", handler.Comment().DeleteSubComment)
		}

	}

}
