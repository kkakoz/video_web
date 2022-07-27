package model

import "video_web/pkg/timex"

type Collection struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Type         uint8      `json:"type"`
	CategoryId   int64      `json:"category_id" gorm:"index"` // 分类
	Cover        string     `json:"cover"`                    // 封面
	Brief        string     `json:"brief"`                    // 简介
	View         int64      `json:"view"`                     // 播放量
	Like         int64      `json:"like"`                     // 喜欢
	Comment      int64      `json:"comment"`                  // 评论数
	Favorite     int64      `json:"favorite"`                 // 收藏
	Duration     int64      `json:"duration"`                 // 时长 /秒
	Hot          int64      `json:"hot"`
	EpisodeCount int64      `json:"episode_count"`
	UserId       int64      `json:"user_id"`
	UserName     string     `json:"user_name"`
	UserAvatar   string     `json:"user_avatar"`
	PublishAt    timex.Time `json:"publish_at"`
	CreatedAt    timex.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    timex.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Orders       string     `json:"orders"`

	User   *User    `json:"user"`
	Videos []*Video `json:"videos" gorm:"-"`
}

type Video struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	CategoryId   int64      `json:"category_id" gorm:"index"` // 分类
	Cover        string     `json:"cover"`                    // 封面
	Brief        string     `json:"brief"`                    // 简介
	UserId       int64      `json:"user_id"`
	UserName     string     `json:"user_name"`
	UserAvatar   string     `json:"user_avatar"`
	Url          string     `json:"url"`
	View         int64      `json:"view"`     // 播放量
	Like         int64      `json:"like"`     // 喜欢
	Comment      int64      `json:"comment"`  // 评论数
	Favorite     int64      `json:"favorite"` // 收藏
	Show         bool       `json:"show"`     // 是否展示在列表
	Duration     int64      `json:"duration"` // 时长 /秒
	CollectionId int64      `json:"collection_id"`
	CreatedAt    timex.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    timex.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Collection *Collection `json:"collection"`
}
