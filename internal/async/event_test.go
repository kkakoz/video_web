package async_test

import (
	"github.com/kkakoz/pkg/redisx"
	"github.com/spf13/viper"
	"testing"
	"time"
	"video_web/internal/async/producer"
	"video_web/internal/model/dto"
	"video_web/pkg/conf"
)

func TestEvents(t *testing.T) {

	conf.InitTestConfig()
	err := redisx.Init(viper.GetViper())
	if err != nil {
		t.Fatal(err)
	}
	//consumer := async.NewEventConsumer(redisx.Client())
	//go consumer.Start(context.Background())

	go func() {
		for {
			err = producer.SendEvent(&dto.Event{
				EventType: dto.EventTypeLike,
				TargetId:  1,
				ActorId:   1,
			})
			if err != nil {
				t.Fatal(err)
			}
			time.Sleep(3 * time.Second)
		}
	}()

	select {}
}
