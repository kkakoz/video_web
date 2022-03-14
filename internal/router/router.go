package router

import (
	"context"
	"net/http"
	"video_web/internal/domain"
	"video_web/pkg/echox"

	"github.com/kkakoz/ormx"
	"github.com/labstack/echo"
	"go.uber.org/fx"
)

func NewHttp(user *userRouter, video *videoRouter, category *categoryRouter) http.Handler {
	e := echo.New()
	e.Binder = echox.NewBinder()
	e.Validator = echox.NewValidator()
	e.HTTPErrorHandler = echox.ErrHandler()
	e.Use(setAccessOriginUrl)
	db := ormx.DB(context.TODO())
	db.AutoMigrate(&domain.User{}, &domain.Auth{}, &domain.Video{}, &domain.Episode{}, &domain.Category{})
	user.AddRouter(e)
	video.AddRouter(e)
	category.AddRouter(e)
	return e
}

var Provider = fx.Provide(NewHttp, NewUserRouter, NewVideoRouter, NewCategoryRouter)
