package bootstrap

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/kkakoz/pkg/app"
	"github.com/kkakoz/pkg/app/kafkas"
	"github.com/kkakoz/pkg/conf"
	"github.com/kkakoz/pkg/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"strings"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/pkg/jsonx"
	"video_web/pkg/ossx"
)

func runVideoHandler() error {
	err := initBase()
	if err != nil {
		return err
	}

	err = ossx.Init(conf.Conf())
	if err != nil {
		return err
	}

	consumer, err := newVideoEventConsumer(conf.Conf())
	if err = app.NewApp("video-web-video-handler", consumer).Start(context.Background()); err != nil {
		return errors.WithMessage(err, "web start err")
	}
	return nil

}

func handler(ctx context.Context, id int64) error {
	bucketPath := "https://kkako-blog-bucket.oss-cn-beijing.aliyuncs.com/"

	video, err := logic.Video().GetById(ctx, id)
	if err != nil {
		return err
	}
	if video == nil {
		return nil
	}
	if video.Cover != "" {
		return nil
	}
	if len(video.Resources) == 0 {
		return nil
	}
	splites := strings.Split(video.Resources[0].Url, "/")
	if len(splites) < 2 {
		return errors.New("http路径错误")
	}
	name := splites[len(splites)-1]
	dirName := splites[len(splites)-2]
	if name == "" || dirName == "" {
		return errors.New("http路径错误")
	}
	coverName := fmt.Sprintf("cover%d.jpg", video.ID)

	currentName := fmt.Sprintf("video%d", video.ID)

	err = exec.CommandContext(ctx, "wget", video.Resources[0].Url, "-O", currentName).Run()
	if err != nil {
		return errors.New("wget error: " + err.Error())
	}

	err = exec.CommandContext(ctx, "ffmpeg", "-ss", "00:00:01", "-i", currentName, "-vframes", "1", coverName).Run()
	if err != nil {
		return errors.New("ffmpeg error: " + err.Error())
	}

	file, err := os.Open(coverName)
	if err != nil {
		return err
	}

	err = ossx.Bucket.PutObject(dirName+"/cover", file)
	if err != nil {
		return err
	}

	return logic.Video().UpdateCover(ctx, video.ID, bucketPath+dirName+"/cover")
}

type EventConsumer struct {
	*kafkas.KafkaConsumer
}

func newVideoEventConsumer(viper *viper.Viper) (*EventConsumer, error) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	consumer, err := kafkas.NewConsumer(viper, eventHandler, config)
	return &EventConsumer{KafkaConsumer: consumer}, err
}

func eventHandler(message *sarama.ConsumerMessage) error {
	event, err := jsonx.Unmarshal[dto.Event](message.Value)
	if err != nil {
		return err
	}
	logger.Info("event = " + string(message.Value))

	if event.EventType == dto.EventTypeAddVideo {
		err := handler(context.Background(), event.TargetId)
		if err != nil {
			logger.Error("设置cover失败", zap.Error(err))
		}
	}

	return nil
}
