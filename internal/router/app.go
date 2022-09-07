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
			app.POST("/resource/get", handler.Resource().Get)
			//app.POST("/resource/list", handler.Resource().GetList)
			app.GET("/resource/ws", handler.Resource().Ws)
		}

		{
			app.GET("/video/get", handler.Video().Get)
			app.POST("/video/list", handler.Video().List)
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

		}

		{
			authApp.POST("/oss/get-conf", handler.Oss().GetConf)
		}

		{
			authApp.POST("/user/current", handler.User().GetCurUser)
		}

		{
			authApp.GET("/video/user-state", handler.Video().UserVideoState)
		}

		{
			authApp.POST("/resource/add", handler.Resource().Add)
			authApp.POST("/resource/del", handler.Resource().Del)
			authApp.POST("/resource/get", handler.Resource().Get)
		}
	}
}
