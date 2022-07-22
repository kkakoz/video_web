package model

import "video_web/pkg/timex"

type Video struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Type         uint8      `json:"type"`
	CategoryId   int64      `json:"category_id" gorm:"index"` // 分类
	Cover        string     `json:"cover"`                    // 封面
	Brief        string     `json:"brief"`                    // 简介
	ViewCount    int64      `json:"view_count"`
	LikeCount    int64      `json:"like_count"`
	Hot          int64      `json:"hot"`
	Collect      int64      `json:"collect"`
	EpisodeCount int64      `json:"episode_count"`
	UserId       int64      `json:"user_id"`
	UserName     string     `json:"user_name"`
	UserAvatar   string     `json:"user_avatar"`
	PublishAt    timex.Time `json:"publish_at"`
	CreatedAt    timex.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    timex.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User     *User      `json:"user"`
	Episodes []*Episode `json:"episodes" gorm:"-"`
}

type Episode struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	UserId int64  `json:"user_id"`
	Url    string `json:"url"`
}

type VideoEpisode struct {
	VideoId   int64 `json:"video_id"`
	EpisodeId int64 `json:"episode_id"`
}
