package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

func BackRouter(e *echo.Echo) {
	back := e.Group("/api/back")

	{
		back.POST("/user/login", handler.User().Login)
	}

	authBack := e.Group("/api/back", authority)

	{

		authBack.POST("/category/list", handler.Category().List)

		{
			// 添加合集
			authBack.POST("/collection/create", handler.Collection().Add)
			authBack.POST("/collection/page-list", handler.Collection().BackList)
		}

		{
			authBack.POST("/video/create", handler.Video().Add)
			// 删除视频
			authBack.POST("/video/del", handler.Video().Del)
			// 视频列表
			authBack.POST("/video/page-list", handler.Video().GetBackList)
		}

		{
			// 当前用户
			authBack.GET("/user/current", handler.User().GetCurUser)
		}

		{
			authBack.POST("/oss/get-conf", handler.Oss().GetConf)
		}

		{
			// 获取评论
			authBack.POST("/comment/get", handler.Comment().Get)
			// 子评论
			authBack.POST("/sub-comment/get", handler.Comment().GetSubComment)
			// 删除评论
			authBack.POST("/comment/del", handler.Comment().Delete)
			authBack.POST("/sub-comment/del", handler.Comment().DeleteSubComment)
		}

	}

}
