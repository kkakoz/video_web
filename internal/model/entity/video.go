package entity

import "video_web/pkg/timex"

type Collection struct {
	ID         int64       `json:"id"`
	Name       string      `json:"name"`
	Type       uint8       `json:"type"`
	CategoryId int64       `json:"category_id" gorm:"index"` // 分类
	Cover      string      `json:"cover"`                    // 封面
	Brief      string      `json:"brief"`                    // 简介
	View       int64       `json:"view"`                     // 播放量
	Like       int64       `json:"like"`                     // 喜欢
	Comment    int64       `json:"comment"`                  // 评论数
	Favorite   int64       `json:"favorite"`                 // 收藏
	Duration   int64       `json:"duration"`                 // 时长 /秒
	Hot        int64       `json:"hot"`
	VideoCount int64       `json:"video_count"`
	UserId     int64       `json:"user_id"`
	UserName   string      `json:"user_name"`
	UserAvatar string      `json:"user_avatar"`
	State      VideoState  `json:"state"`
	PublishAt  *timex.Time `json:"publish_at"`
	CreatedAt  timex.Time  `json:"created_at"`
	UpdatedAt  timex.Time  `json:"updated_at"`
	Orders     string      `json:"orders"`

	User   *User    `json:"user"`
	Videos []*Video `json:"videos"`
}

type CollectionType uint8

const (
	CollectionTypeAnime    CollectionType = 1
	CollectionTypeSeries   CollectionType = 2
	CollectionTypeDiv      CollectionType = 3
	CollectionTypeTransfer CollectionType = 4
)

type Video struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Type         VideoType  `json:"type"`
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
	Duration     int64      `json:"duration"` // 时长 /秒
	CollectionId int64      `json:"collection_id"`
	State        VideoState `json:"state"` //
	CreatedAt    timex.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    timex.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Collection *Collection `json:"collection"`
}

type VideoType uint8

const (
	VideoTypeDiv      VideoType = 1 // 自制
	VideoTypeTransfer VideoType = 2 // 转载
)

type VideoState uint8

const (
	VideoStateNormal   VideoState = 1 // 正常展示
	VideoStateHidden   VideoState = 2 // 隐藏
	VideoStateUserBan  VideoState = 3 // 禁止
	VideoStateAdminBan VideoState = 4 // admin禁止
)
