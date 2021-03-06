package router

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"net/http"
	"video_web/internal/model"
	"video_web/pkg/echox"

	"github.com/labstack/echo"
	"go.uber.org/fx"
)

func NewHttp(user *userRouter, video *videoRouter, category *categoryRouter,
	logger *zap.Logger, comment *commentRouter, like *likeRouter, follow *followRouter, oss *ossRouter) http.Handler {
	e := echo.New()
	e.Binder = echox.NewBinder()
	e.Validator = echox.NewValidator()
	e.HTTPErrorHandler = echox.ErrHandler(logger)
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	// e.Use(setAccessOriginUrl)
	db := ormx.DB(context.TODO())
	// db.AutoMigrate(&model.Follow{}, &model.FollowGroup{})
	db.AutoMigrate(&model.User{}, &model.Video{}, &model.Episode{},
		&model.Category{}, &model.Comment{}, &model.SubComment{}, &model.VideoEpisode{}, &model.UserSecurity{})
	e.Debug = true
	user.AddRouter(e)
	video.AddRouter(e)
	category.AddRouter(e)
	comment.AddRouter(e)
	like.AddRouter(e)
	follow.AddRouter(e)
	oss.AddRouter(e)
	return e
}

var Provider = fx.Provide(NewHttp, NewUserRouter, NewVideoRouter, NewCategoryRouter, NewCommentRouter, NewLikeRouter, NewFollowRouter, NewOssRouter)
