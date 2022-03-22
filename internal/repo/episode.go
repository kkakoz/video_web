package repo

import (
	"video_web/internal/domain"

	"github.com/kkakoz/ormx"
)

var _ domain.IEpisodeRepo = (*EpisodeRepo)(nil)

type EpisodeRepo struct {
	ormx.IRepo[domain.Episode]
}

func NewEpisodeRepo() domain.IEpisodeRepo {
	return &EpisodeRepo{
		ormx.NewRepo[domain.Episode](),
	}
}
