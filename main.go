package main

import (
	"github.com/kkakoz/ormx"
	"github.com/pkg/errors"
	"log"
	"video_web/internal/router"
	"video_web/pkg/app"
	"video_web/pkg/conf"
)

func main() {

	conf.InitConfig()
	if _, err := ormx.New(conf.Conf()); err != nil {
		log.Fatalln("init mysql conn err:", err)
	}
	ormx.DefaultErrHandler = func(err error) error {
		return errors.WithStack(err)
	}

	if err := app.NewApplication("video_web", router.NewHttp(), []app.Server{}).Run(); err != nil {
		log.Fatalln(err)
	}
}
