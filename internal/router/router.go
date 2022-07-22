package router

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/labstack/echo/middleware"
	"net/http"
	"video_web/internal/model"
	"video_web/pkg/echox"
	"video_web/pkg/logger"

	"github.com/labstack/echo"
)

func NewHttp() http.Handler {
	e := echo.New()
	e.Binder = echox.NewBinder()
	e.Validator = echox.NewValidator()
	e.HTTPErrorHandler = echox.ErrHandler(logger.L())
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	// e.Use(setAccessOriginUrl)
	db := ormx.DB(context.TODO())
	// db.AutoMigrate(&model.Follow{}, &model.FollowGroup{})
	db.AutoMigrate(&model.User{}, &model.Video{}, &model.Episode{}, &model.FollowGroup{}, &model.Follow{},
		&model.Category{}, &model.Comment{}, &model.SubComment{}, &model.VideoEpisode{}, &model.UserSecurity{})

	e.Debug = true
	AppRouter(e)
	BackRouter(e)
	return e
}
