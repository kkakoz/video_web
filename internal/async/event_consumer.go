package async

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/kkakoz/pkg/app/kafkas"
	"github.com/kkakoz/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"video_web/internal/model/dto"
	"video_web/pkg/jsonx"
)

type EventConsumer struct {
	*kafkas.KafkaConsumer
}

func NewEventConsumer(viper *viper.Viper) (*EventConsumer, error) {
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

	logger.Debug("consumer event:" + string(message.Value))

	eventHandler := handlerMap[event.EventType]
	err = eventHandler(event)
	if err != nil {
		logger.Error("consumer event err", zap.Error(err), zap.String("event", string(message.Value)))
	}
	return nil
}
