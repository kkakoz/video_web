package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

func BackRouter(e *echo.Echo) {

	back := e.Group("/api/back", setUser)
	{

		back.POST("/user/login", handler.User().Login)
	}

	authBack := back.Group("", authority)

	{

		authBack.GET("/category/list", handler.Category().List)

		{
			// 添加合集
			authBack.POST("/video/add", handler.Video().Add)
			authBack.GET("/video/page-list", handler.Video().BackList)
		}

		{
			authBack.GET("/uv/calculate", handler.Statistics().CalculateUV)
			authBack.GET("/dau/calculate", handler.Statistics().CalculateDAU)
		}

		{
			// 添加稿件
			authBack.POST("/resource/add", handler.Resource().Add)
			// 添加分p视频
			authBack.POST("/resource/add-list", handler.Resource().AddList)
			// 删除视频
			authBack.POST("/resource/del", handler.Resource().Del)
			// 视频列表
			authBack.GET("/resource/page-list", handler.Resource().GetBackList)
		}

		{
			// 当前用户
			authBack.GET("/user/current", handler.User().GetCurUser)
		}

		{
			authBack.GET("/oss/get-conf", handler.Oss().GetConf)
		}

		{
			// 获取评论
			authBack.GET("/comment/get", handler.Comment().List)
			// 子评论
			authBack.GET("/sub-comment/get", handler.Comment().GetSubComment)
			// 删除评论
			authBack.POST("/comment/del", handler.Comment().Delete)
			authBack.POST("/sub-comment/del", handler.Comment().DeleteSubComment)
		}

	}

}
