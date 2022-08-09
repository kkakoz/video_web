package main

import (
	"github.com/kkakoz/ormx"
	"github.com/pkg/errors"
	"log"
	"time"
	"video_web/internal/router"
	"video_web/pkg/app"
	"video_web/pkg/conf"
	"video_web/pkg/redisx"
)

func main() {

	var LOC, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic("初始化时区失败")
	}
	time.Local = LOC

	conf.InitConfig()
	if _, err := ormx.New(conf.Conf()); err != nil {
		log.Fatalln("init mysql conn err:", err)
	}
	ormx.DefaultErrHandler = func(err error) error {
		return errors.WithStack(err)
	}
	err = redisx.Init(conf.Conf())
	if err != nil {
		log.Fatalln(err)
	}

	if err := app.NewApplication("video_web", router.NewHttp(), []app.Server{}).Run(); err != nil {
		log.Fatalln(err)
	}
}
