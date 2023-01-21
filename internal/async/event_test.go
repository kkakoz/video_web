package async_test

import (
	"github.com/kkakoz/pkg/conf"
	"github.com/kkakoz/pkg/redisx"
	"testing"
	"time"
	"video_web/internal/async/producer"
	"video_web/internal/model/dto"
)

func TestEvents(t *testing.T) {

	err := redisx.Init(conf.Conf())
	if err != nil {
		t.Fatal(err)
	}
	//consumer := async.NewEventConsumer(redisx.Client())
	//go consumer.Start(context.Background())

	go func() {
		for {
			err = producer.Send(&dto.Event{
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
