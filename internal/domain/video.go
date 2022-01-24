package domain

import (
	"context"
	"video_web/pkg/gormx"
	"video_web/pkg/model"
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
	CreatedAt  int64      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  int64      `json:"updated_at" gorm:"autoUpdateTime"`
	User       *User      `json:"user"`
}

const (
	VideoTypeNormal = 1
)

type Episode struct {
	ID      int64  `json:"id"`
	VideoId int64  `json:"video_id"`
	Order   int64  `json:"order"`
	Url     string `json:"url"`
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
	GetVideo(ctx context.Context, opts ...gormx.DBOption) (*Video, error)
	GetVideos(ctx context.Context, video *Video, pager *model.Pager) ([]*Video, error)

	AddEpisode(ctx context.Context, episode *Episode) error
	GetEpisode(ctx context.Context, episodeId int64) (*Episode, error)
	GetEpisodes(ctx context.Context, videoId int64) ([]*Episode, error)
	UpdateEpisode(ctx context.Context, episode *Episode) error
	DelEpisode(ctx context.Context, episodeId int64) error
	UpdateAfterOrderEpisode(ctx context.Context, videoId int64, order int64, updateVal int) error
}
