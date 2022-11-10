package router

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/labstack/echo/middleware"
	"net/http"
	"video_web/internal/model/entity"
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
	db.AutoMigrate(&entity.User{}, &entity.Video{}, &entity.Resource{}, &entity.FollowGroup{}, &entity.Follow{}, &entity.Collect{}, entity.CollectGroup{},
		&entity.Category{}, &entity.Comment{}, &entity.SubComment{}, &entity.UserSecurity{}, &entity.Like{}, &entity.Newsfeed{}, &entity.History{})

	e.Debug = true
	AppRouter(e)
	BackRouter(e)
	return e
}
