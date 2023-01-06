package jobs

import (
	"context"
	"github.com/kkakoz/pkg/logger"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"video_web/internal/logic"
)

type Jobs struct {
	c *cron.Cron
}

func (j *Jobs) Start(ctx context.Context) error {
	j.c = cron.New()

	_, err := j.c.AddFunc("*/15 * * * *", func() {
		err := logic.Video().CalculateHot(context.Background())
		if err != nil {
			logger.Error("Error calculate hot", zap.Error(err))
		}
	})
	if err != nil {
		return err
	}

	j.c.Start()
	select {
	case <-ctx.Done():
		return nil
	}
}

func (j *Jobs) Stop(ctx context.Context) error {
	<-j.c.Stop().Done()
	return nil
}
