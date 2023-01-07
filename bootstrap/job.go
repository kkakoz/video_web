package bootstrap

import (
	"context"
	"github.com/kkakoz/pkg/app"
	"github.com/pkg/errors"
	"video_web/internal/job"
)

// job任务执行
func runJob() error {
	err := initBase()
	if err != nil {
		return err
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
