package bootstrap

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/pkg/app"
	"github.com/kkakoz/pkg/app/https"
	"github.com/kkakoz/pkg/logger"
	"github.com/kkakoz/pkg/redisx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"video_web/internal/async"
	"video_web/internal/async/producer"
	"video_web/internal/router"
	"video_web/pkg/conf"
)

// web程序
func run() error {
	logger.InitLog(viper.GetViper())
	if _, err := ormx.New(conf.Conf()); err != nil {
		return errors.WithMessage(err, "init orm failed")
	}
	ormx.DefaultErrHandler = func(err error) error {
		return errors.WithStack(err)
	}
	err := redisx.Init(conf.Conf())
	if err != nil {
		return errors.WithMessage(err, "init redis failed")
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
