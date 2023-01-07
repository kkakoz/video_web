package bootstrap

import (
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/pkg/logger"
	"github.com/kkakoz/pkg/redisx"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"time"
	"video_web/pkg/conf"
)

var cfg = pflag.StringP("config", "c", "configs/conf.yaml", "Configuration file.")

func Run() error {
	time.Local = time.FixedZone("UTC+8", 8*3600)
	pflag.Parse()
	conf.InitConfig(*cfg)

	arg0 := pflag.Arg(0)

	switch arg0 {
	case "migrate":
		return migrate()
	case "job":
		return runJob()
	default:
		return run()
	}

}

func initBase() error {
	logger.InitLog(conf.Conf())
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
	return nil
}
