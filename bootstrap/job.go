package bootstrap

import (
	"context"
	"github.com/kkakoz/ormx"
	"github.com/kkakoz/pkg/app"
	"github.com/kkakoz/pkg/logger"
	"github.com/kkakoz/pkg/redisx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"video_web/internal/job"
	"video_web/pkg/conf"
)

// job任务执行
func runJob() error {
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

	jobs := job.NewJobs()

	servers := []app.Server{
		jobs,
	}

	if err = app.NewApp("video-web-job", servers...).Start(context.Background()); err != nil {
		return errors.WithMessage(err, "web start err")
	}
	return nil
}
