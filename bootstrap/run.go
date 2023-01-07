package bootstrap

import (
	"context"
	"github.com/kkakoz/pkg/app"
	"github.com/kkakoz/pkg/app/https"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"video_web/internal/async"
	"video_web/internal/async/producer"
	"video_web/internal/router"
)

// web程序
func run() error {
	err := initBase()
	if err != nil {
		return err
	}

	err = producer.Init(viper.GetViper())
	if err != nil {
		return errors.WithMessage(err, "init kafka producer failed")
	}

	eventConsumer, err := async.NewEventConsumer(viper.GetViper())
	if err != nil {
		return errors.WithMessage(err, "event consumer init err")
	}

	servers := []app.Server{
		https.NewHttpServer(router.NewHttp(), ":"+viper.GetString("app.port")),
		eventConsumer,
	}

	if err = app.NewApp("video_web", servers...).Start(context.Background()); err != nil {
		return errors.WithMessage(err, "web start err")
	}
	return nil
}
