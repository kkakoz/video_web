package repo

import (
	"sync"
	"video_web/internal/model"

	"github.com/kkakoz/ormx"
)

type videoRepo struct {
	ormx.IRepo[model.Video]
}

var videoOnce sync.Once
var _video *videoRepo

func Video() *videoRepo {
	videoOnce.Do(func() {
		_video = &videoRepo{IRepo: ormx.NewRepo[model.Video]()}
	})
	return _video
}
