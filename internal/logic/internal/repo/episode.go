package repo

import (
	"sync"
	"video_web/internal/model"

	"github.com/kkakoz/ormx"
)

type episodeRepo struct {
	ormx.IRepo[model.Episode]
}

var episodeOnce sync.Once
var _episode *episodeRepo

func Episode() *episodeRepo {
	episodeOnce.Do(func() {
		_episode = &episodeRepo{IRepo: ormx.NewRepo[model.Episode]()}
	})
	return _episode
}
