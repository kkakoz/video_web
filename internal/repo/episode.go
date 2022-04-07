package repo

import (
	"video_web/internal/model"

	"github.com/kkakoz/ormx"
)

type EpisodeRepo struct {
	ormx.IRepo[model.Episode]
}

func NewEpisodeRepo() *EpisodeRepo {
	return &EpisodeRepo{
		ormx.NewRepo[model.Episode](),
	}
}
