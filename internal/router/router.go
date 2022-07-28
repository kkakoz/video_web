package router

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/labstack/echo/middleware"
	"net/http"
	entity2 "video_web/internal/model/entity"
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
	db.AutoMigrate(&entity2.User{}, &entity2.Collection{}, &entity2.Video{}, &entity2.FollowGroup{}, &entity2.Follow{},
		&entity2.Category{}, &entity2.Comment{}, &entity2.SubComment{}, &entity2.UserSecurity{})

	e.Debug = true
	AppRouter(e)
	BackRouter(e)
	return e
}
