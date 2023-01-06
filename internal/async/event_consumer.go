package async

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/kkakoz/pkg/logger"
	"go.uber.org/zap"
	"video_web/internal/model/dto"
	"video_web/pkg/jsonx"
)

type EventConsumer struct {
	cli *redis.Client
}

func NewEventConsumer(cli *redis.Client) *EventConsumer {
	return &EventConsumer{cli: cli}
}

func (e *EventConsumer) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			result, err := e.cli.BRPop(0, dto.EventKey).Result()
			if err != nil {
				logger.Error("brpop consumer event failed", zap.Error(err))
				continue
			}
			fmt.Println("result = ", result)
			for i, res := range result {
				if i == 0 {
					continue
				}
				event, err := jsonx.Unmarshal[dto.Event]([]byte(res))
				if err != nil {
					logger.L().Error("unmarshal consumer event failed", zap.Error(err))
				}
				fmt.Println("event = ", event)
				eventHandler := handlerMap[event.EventType]
				err = eventHandler(event)
				if err != nil {
					logger.Error("consumer event err", zap.Error(err), zap.String("event", res))
					continue
				}

			}
		}
	}
}

func (e *EventConsumer) Stop(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
