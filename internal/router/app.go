package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

func AppRouter(e *echo.Echo) {
	app := e.Group("/api/app")

	{

		{
			app.GET("/category/list", handler.Category().List)
		}

		{
			app.POST("/user/login", handler.User().Login)
			app.GET("/user", handler.User().GetUser)
			app.POST("/user/create", handler.User().Register)
		}

		{
			app.GET("/video", handler.Video().Get)
			app.GET("/video/list", handler.Video().GetList)
			app.GET("/video/ws", handler.Video().Ws)
		}

		{
			app.GET("/collection", handler.Collection().Get)
			app.GET("/collection/list", handler.Collection().List)
		}

	}

	authApp := e.Group("/api/app", authority)

	{

		{
			authApp.POST("/comments/create", handler.Comment().Add)
			authApp.POST("/sub-comments/create", handler.Comment().AddSubComment)
			authApp.GET("/comments", handler.Comment().Get)
			authApp.GET("/sub-comments/list", handler.Comment().GetSubComment)
		}

		{
			authApp.POST("/follow", handler.Follow().Follow)
			authApp.GET("/follow/fans", handler.Follow().Fans)
			authApp.GET("/followers/list", handler.Follow().Followers)
			authApp.GET("/followed/list", handler.Follow().IsFollowed)
			authApp.POST("/follow-groups", handler.Follow().GroupAdd)
			authApp.GET("/follow-groups/list", handler.Follow().Groups)
		}

		{
			authApp.POST("/like/create", handler.Like().Like)
			authApp.GET("/like", handler.Like().IsLike)
		}

		{
			authApp.GET("/oss/conf", handler.Oss().GetConf)
		}

		{
			authApp.GET("/user/current", handler.User().GetCurUser)
		}

		{
			//authApp.POST("/video/create", handler.Video().Add)
			authApp.POST("/video/del", handler.Video().Del)
			authApp.GET("/video", handler.Video().Get)
			authApp.GET("/video/list", handler.Video().GetList)
		}
	}
}
