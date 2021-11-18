package main

import (
	"github.com/spf13/viper"
	"github/kkakoz/video_web/internal/handler"
	"github/kkakoz/video_web/internal/router"
	"github/kkakoz/video_web/internal/server"
	"github/kkakoz/video_web/pkg/app"
	"github/kkakoz/video_web/pkg/conf"
	"github/kkakoz/video_web/pkg/mysqlx"
	"github/kkakoz/video_web/pkg/redis"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func NewApp(handler http.Handler, servers []app.Server) *app.Application {
	return app.NewApplication("video_web", handler, servers)
}

func main() {

	conf.InitConfig()
	if _, err := mysqlx.New(viper.GetViper()); err != nil {
		log.Fatalln("init mysql conn err:", err)
	}

	var app = new(app.Application)
	fx.New(
		handler.Provider,
		redis.Provider,
		server.Provider,
		router.Provider,
		fx.Provide(NewApp),
		fx.Supply(viper.GetViper()),
		fx.Populate(&app),
	)
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
