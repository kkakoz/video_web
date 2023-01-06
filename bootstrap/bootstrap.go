package bootstrap

import (
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/pkg/logger"
	"github.com/kkakoz/pkg/redisx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"time"
	"video_web/pkg/conf"
)

func BootstrapInit() error {
	time.Local = time.FixedZone("UTC+8", 8*3600)

	conf.InitConfig()
	if _, err := ormx.New(conf.Conf()); err != nil {
		log.Fatalln("init mysql conn err:", err)
	}
	ormx.DefaultErrHandler = func(err error) error {
		return errors.WithStack(err)
	}
	err := redisx.Init(conf.Conf())
	if err != nil {
		return err
	}

	logger.InitLog(viper.GetViper())

	return nil

}
