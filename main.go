package main

import (
	"context"
	"github.com/kkakoz/pkg/app"
	"github.com/kkakoz/pkg/app/https"
	"github.com/kkakoz/pkg/logger"
	"github.com/kkakoz/pkg/redisx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"video_web/bootstrap"
	"video_web/internal/async"
	"video_web/internal/router"
)

func main() {

	err := bootstrap.BootstrapInit()
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}

	servers := []app.Server{
		https.NewHttpServer(router.NewHttp(), ":"+viper.GetString("app.port")),
		async.NewEventConsumer(redisx.Client()),
	}

	if err = app.NewApp("video_web", servers...).Start(context.Background()); err != nil {
		log.Fatalln(err)
	}
}
