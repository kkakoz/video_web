package domain

import (
	"context"
	"github.com/kkakoz/ormx"
	"video_web/internal/dto/request"
)

type Video struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Type         uint8      `json:"type"`
	CategoryId   int64      `json:"category_id" gorm:"index"` // 分类
	Cover        string     `json:"cover"`                    // 封面
	Brief        string     `json:"brief"`
	ViewCount    int64      `json:"view_count"`
	LikeCount    int64      `json:"like_count"`
	Hot          int64      `json:"hot"`
	Collect      int64      `json:"collect"`
	EpisodeCount int64      `json:"episode_count"`
	Episodes     []*Episode `json:"episodes"`
	UserId       int64      `json:"user_id"`
	UserName     string     `json:"user_name"`
	UserAvatar   string     `json:"user_avatar"`
	CreatedAt    int64      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    int64      `json:"updated_at" gorm:"autoUpdateTime"`
	User         *User      `json:"user"`
}

type Episode struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	UserId int64  `json:"user_id"`
	Order  int64  `json:"order"`
	Url    string `json:"url"`
	Video  *Video `json:"video"`
}

type VideoEpisode struct {
	VideoId   int64 `json:"video_id"`
	EpisodeId int64 `json:"episode_id"`
}

type IVideoLogic interface {
	Add(ctx context.Context, video *request.AddVideoReq) error
	GetVideo(ctx context.Context, videoId int64) (*Video, error)
	GetVideos(ctx context.Context, categoryId uint, lastId uint, orderType uint8) ([]*Video, error)

	AddEpisode(ctx context.Context, req *request.AddEpisodeReq) error
	DelVideoEpisode(ctx context.Context, req *request.EpisodeIdReq) error
}

type IVideoRepo interface {
	ormx.IRepo[Video]
	UpdateAfterOrderEpisode(ctx context.Context, videoId int64, order int64, updateVal int) error
	AddEpisode(ctx context.Context, videoId int64, episodeIds []int64) error
	GetEpisodeIds(ctx context.Context, videoId int64) ([]int64, error)
	DeleteEpisode(ctx context.Context, videoId int64, episodeId int64) error
}

type IEpisodeRepo interface {
	ormx.IRepo[Episode]
}
