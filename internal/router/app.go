package router

import (
	"github.com/labstack/echo"
	"video_web/internal/handler"
)

func AppRouter(e *echo.Echo) {
	app := e.Group("/api/app")

	{
		categoryG := app.Group("/categories")
		{
			categoryG.GET("", handler.Category().List)
		}

		commentG := app.Group("/comments")
		{
			commentG.GET("", handler.Comment().Get)
			commentG.GET("/:comment_id/sub_comments", handler.Comment().GetSubComment)
		}

		userG := app.Group("/users")
		{
			userG.POST("", handler.User().Register)
			userG.POST("/login", handler.User().Login)
			userG.GET("/:user_id", handler.User().GetUser)
		}

		videoG := app.Group("/videos")
		{
			videoG.GET("/:video_id", handler.Video().Get)
			videoG.GET("", handler.Video().GetList)
			videoG.GET("/:video_id/ws", handler.Video().Ws)
		}

		collectionG := app.Group("/collections")
		{
			collectionG.GET("/:collection_id", handler.Collection().Get)
		}

	}

	authApp := e.Group("/api/app", authority)

	{
		commentG := authApp.Group("/comments")
		{
			commentG.POST("", handler.Comment().Add)
			commentG.POST("/:comment_id/sub_comments", handler.Comment().AddSubComment)
		}

		followG := authApp.Group("/follow")
		{
			followG.POST("", handler.Follow().Follow)
			followG.GET("/fans", handler.Follow().Fans)
			followG.GET("/followers", handler.Follow().Followers)
			followG.GET("/followed", handler.Follow().IsFollowed)
			followG.POST("/groups", handler.Follow().GroupAdd)
			followG.GET("/groups", handler.Follow().Groups)
		}

		likeG := authApp.Group("/likes")
		{
			likeG.POST("", handler.Like().Like)
			likeG.GET("", handler.Like().IsLike)
		}

		OssG := authApp.Group("/oss")

		OssG.GET("/conf", handler.Oss().GetConf)

		userG := authApp.Group("/users")
		{
			userG.GET("/local", handler.User().GetCurUser)
		}

		videoG := authApp.Group("/videos")
		{
			videoG.POST("", handler.Video().Add)
			videoG.DELETE("/:video_id", handler.Video().Del)
		}
	}
}
