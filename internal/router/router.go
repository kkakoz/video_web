package router

import (
	"context"
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/internal/handler"
	"github/kkakoz/video_web/pkg/echox"
	"github/kkakoz/video_web/pkg/mysqlx"
	"go.uber.org/fx"
	"net/http"
)

func NewHttp(user *handler.UserHandler) http.Handler {
	e := echo.New()
	e.Binder = echox.NewBinder()
	e.Validator = echox.NewValidator()
	e.HTTPErrorHandler = echox.ErrHandler()
	db := mysqlx.GetDB(context.TODO())
	db.AutoMigrate(&domain.User{}, &domain.Auth{})
	userRouter(e, user)
	return e
}

var Provider = fx.Provide(NewHttp)