package event_send

import (
	"encoding/json"
	"github.com/kkakoz/pkg/redisx"
	"video_web/internal/model/dto"
)

func SendEvent(event *dto.Event) error {
	cli := redisx.Client()

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return cli.LPush(dto.EventKey, data).Err()
}
