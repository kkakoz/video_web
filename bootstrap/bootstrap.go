package bootstrap

import (
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/pkg/conf"
	"github.com/kkakoz/pkg/redisx"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"time"
	"video_web/pkg/logs"
)

func Run() error {
	time.Local = time.FixedZone("UTC+8", 8*3600)
	pflag.Parse()

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
	if _, err := ormx.New(conf.Conf(), ormx.WithLogger(logs.NewGLogger())); err != nil {
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
