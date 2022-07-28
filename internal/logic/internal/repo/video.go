package repo

import (
	"sync"
	"video_web/internal/model/entity"

	"github.com/kkakoz/ormx"
)

type videoRepo struct {
	ormx.IRepo[entity.Video]
}

var videoOnce sync.Once
var _video *videoRepo

func Video() *videoRepo {
	videoOnce.Do(func() {
		_video = &videoRepo{IRepo: ormx.NewRepo[entity.Video]()}
	})
	return _video
}
