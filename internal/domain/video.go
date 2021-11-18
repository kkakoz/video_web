package domain

import (
	"context"
	"github/kkakoz/video_web/pkg/model"
)

type Video struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Type       int32      `json:"type"`
	Category   int32      `json:"category"` // 分类
	Cover      string     `json:"cover"`    // 封面
	Brief      string     `json:"brief"`
	View       int64      `json:"view"`
	Like       int64      `json:"like"`
	Collect    int64      `json:"collect"`
	Episodes   []*Episode `json:"episodes"`
	UserId     int64      `json:"user_id"`
	UserName   string     `json:"user_name"`
	UserAvatar string     `json:"user_avatar"`
	CreatedAt  int64      `gorm:"autoCreateTime"`
	UpdatedAt  int64      `gorm:"autoUpdateTime"`
}

type Episode struct {
	ID      int64 `json:"id"`
	VideoId int64 `json:"video_id"`
	PreId   int64 `json:"pre_id"`
	NextId  int64 `json:"next_id"`
	Url     int64 `json:"url"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type IVideoLogic interface {
	AddVideo(ctx context.Context, video *Video) error
	GetVideo(ctx context.Context, videoId int64) (*Video, error)
	GetVideos(ctx context.Context, video *Video, pager *model.Pager) ([]*Video, error)

	AddEpisode(ctx context.Context, episode *Episode) error
	DelEpisode(ctx context.Context, episodeId int64) error
}

type IVideoRepo interface {
	AddVideo(ctx context.Context, video *Video) error
	GetVideo(ctx context.Context, videoId int64) (*Video, error)
	GetVideos(ctx context.Context, video *Video, pager *model.Pager) ([]*Video, error)

	AddEpisode(ctx context.Context, episode *Episode) error
	GetEpisode(ctx context.Context, episodeId int64) (*Episode, error)
	GetEpisodes(ctx context.Context, videoId int64) ([]*Episode, error)
	GetLastEpisode(ctx context.Context, videoId int64) (*Episode, error)
	UpdateEpisode(ctx context.Context, episode *Episode) error
	DelEpisode(ctx context.Context, episodeId int64) error
}
