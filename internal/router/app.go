package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

func AppRouter(e *echo.Echo) {
	app := e.Group("/api/app")

	{

		{
			app.POST("/category/list", handler.Category().List)
		}

		{
			app.POST("/user/login", handler.User().Login)
			app.POST("/user/get", handler.User().GetUser)
			app.POST("/user/create", handler.User().Register)
		}

		{
			app.POST("/video/get", handler.Video().Get)
			app.POST("/video/list", handler.Video().GetList)
			app.GET("/video/ws", handler.Video().Ws)
		}

		{
			app.POST("/collection/get", handler.Collection().Get)
		}

	}

	authApp := e.Group("/api/app", authority)

	{

		{
			authApp.POST("/comments/create", handler.Comment().Add)
			authApp.POST("/sub-comments/create", handler.Comment().AddSubComment)
			authApp.POST("/comments/get", handler.Comment().Get)
			authApp.POST("/sub-comments/list", handler.Comment().GetSubComment)
		}

		{
			authApp.POST("/follow", handler.Follow().Follow)
			authApp.POST("/follow/fans", handler.Follow().Fans)
			authApp.POST("/followers/list", handler.Follow().Followers)
			authApp.POST("/followed/list", handler.Follow().IsFollowed)
			authApp.POST("/follow-groups", handler.Follow().GroupAdd)
			authApp.POST("/follow-groups/list", handler.Follow().Groups)
		}

		{
			authApp.POST("/like/create", handler.Like().Like)
			authApp.POST("/like/get", handler.Like().IsLike)
		}

		{
			authApp.POST("/oss/get-conf", handler.Oss().GetConf)
		}

		{
			authApp.POST("/user/current", handler.User().GetCurUser)
		}

		{
			authApp.POST("/video/create", handler.Video().Add)
			authApp.POST("/video/del", handler.Video().Del)
			authApp.POST("/video/get", handler.Video().Get)
			authApp.POST("/video/list", handler.Video().GetList)
		}
	}
}
