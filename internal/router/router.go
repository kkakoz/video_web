package router

import (
	"context"
	"github.com/kkakoz/ormx"
	"go.uber.org/zap"
	"net/http"
	"video_web/internal/domain"
	"video_web/pkg/echox"

	"github.com/labstack/echo"
	"go.uber.org/fx"
)

func NewHttp(user *userRouter, video *videoRouter, category *categoryRouter, logger *zap.Logger, comment *commentRouter, like *likeRouter) http.Handler {
	e := echo.New()
	e.Binder = echox.NewBinder()
	e.Validator = echox.NewValidator()
	e.HTTPErrorHandler = echox.ErrHandler(logger)
	e.Use(setAccessOriginUrl)
	db := ormx.DB(context.TODO())
	db.AutoMigrate(&domain.Like{})
	// db.AutoMigrate(&domain.User{}, &domain.Auth{}, &domain.Video{}, &domain.Episode{},
	// 	&domain.Category{}, &domain.Comment{}, &domain.SubComment{})
	e.Debug = true
	user.AddRouter(e)
	video.AddRouter(e)
	category.AddRouter(e)
	comment.AddRouter(e)
	like.AddRouter(e)
	return e
}

var Provider = fx.Provide(NewHttp, NewUserRouter, NewVideoRouter, NewCategoryRouter, NewCommentRouter, NewLikeRouter)
