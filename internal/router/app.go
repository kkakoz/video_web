package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

func AppRouter(e *echo.Echo) {

	app := e.Group("/api/app", setUser)

	{

		{
			app.GET("/category/list", handler.Category().List)
		}

		{
			app.POST("/user/login", handler.User().Login)
			app.GET("/user/get", handler.User().GetUser)
			app.POST("/user/register", handler.User().Register)
		}

		{
			app.GET("/resource/list", handler.Resource().List)
			//app.POST("/resource/list", handler.Resource().GetList)
			app.GET("/resource/ws", handler.Resource().Ws)
		}

		{
			app.GET("/video/get", handler.Video().Get)
			app.GET("/video/list", handler.Video().List)
			app.GET("/video/recommend", handler.Video().Recommend)
			app.GET("/video/rankings", handler.Video().Rankings)
		}

		{
			app.GET("/comment/list", handler.Comment().List)
			app.GET("/sub-comment/list", handler.Comment().GetSubComment)
		}

		{
			app.GET("/newsfeed/page-list", handler.Newsfeed().UserNews)
		}

	}

	{
		authApp := app.Group("", authority)

		{
			authApp.POST("/comment/add", handler.Comment().Add)
			authApp.POST("/sub-comment/add", handler.Comment().AddSubComment)
		}

		{
			authApp.POST("/user/update-avatar", handler.User().UpdateAvatar)
		}

		{
			authApp.POST("/follow/add", handler.Follow().Follow)
			authApp.GET("/follow/fans", handler.Follow().Fans)
			authApp.GET("/follower/list", handler.Follow().Followers)
			authApp.GET("/followed/list", handler.Follow().IsFollowed)
			authApp.POST("/follow-group/add", handler.Follow().GroupAdd)
			authApp.GET("/follow-group/list", handler.Follow().Groups)
		}

		{
			authApp.POST("/like/add", handler.Like().Like)

		}

		{
			authApp.POST("/oss/get-conf", handler.Oss().GetConf)
		}

		{
			authApp.GET("/video/user-state", handler.Video().UserVideoState)
		}

		{
			authApp.GET("/user/current", handler.User().GetCurUser)
		}

		{
			authApp.GET("/oss/get-conf", handler.Oss().GetConf)
		}

		{
			authApp.POST("/resource/add", handler.Resource().Add)
			authApp.POST("/resource/del", handler.Resource().Del)
			authApp.POST("/resource/get", handler.Resource().Get)
		}

		{
			authApp.POST("/collect/add", handler.Collect().Add)
			authApp.GET("/collect/page-list", handler.Collect().List)
		}

		{
			authApp.POST("/newsfeed/add", handler.Newsfeed().Add)
			authApp.GET("/newsfeed/followed-page-list", handler.Newsfeed().FollowedNewsFeeds)
		}

		{
			authApp.POST("/history/add", handler.History().Add)
			authApp.GET("/history/page-list", handler.History().List)
			authApp.GET("/history/get", handler.History().Get)
		}
	}
}
