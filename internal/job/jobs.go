package job

import (
	"context"
	"github.com/kkakoz/pkg/logger"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"video_web/internal/logic"
)

type jobs struct {
	c *cron.Cron
}

func NewJobs() *jobs {
	c := cron.New()
	return &jobs{c: c}
}

func (j *jobs) Start(ctx context.Context) error {

	calculateHot()
	_, err := j.c.AddFunc("*/15 * * * *", calculateHot)
	if err != nil {
		return err
	}
	updateLike()
	_, err = j.c.AddFunc("*/15 * * * *", updateLike)
	if err != nil {
		return err
	}

	j.c.Start()
	select {
	case <-ctx.Done():
		return nil
	}
}

func (j *jobs) Stop(ctx context.Context) error {
	<-j.c.Stop().Done()
	return nil
}

func calculateHot() {
	logger.Info("计算hot")
	err := logic.Video().CalculateHot(context.Background())
	if err != nil {
		logger.Error("Error calculate hot", zap.Error(err))
	}
}

func updateLike() {
	logger.Info("同步缓存like")
	err := logic.Like().UpdateLikeJob(context.Background())
	if err != nil {
		logger.Error("Error update like count", zap.Error(err))
	}
}
