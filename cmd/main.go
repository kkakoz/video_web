package main

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"net/http"
	"video_web/internal/handler"
	"video_web/internal/router"
	"video_web/internal/server"
	"video_web/pkg/app"
	"video_web/pkg/conf"
	"video_web/pkg/mysqlx"
	"video_web/pkg/redis"
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
