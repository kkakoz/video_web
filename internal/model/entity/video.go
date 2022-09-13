package entity

import "video_web/pkg/timex"

type Collection struct {
	ID         int64           `json:"id"`
	Name       string          `json:"name"`
	Type       uint8           `json:"type"`
	CategoryId int64           `json:"category_id" gorm:"index"` // 分类
	Cover      string          `json:"cover"`                    // 封面
	Brief      string          `json:"brief"`                    // 简介
	View       int64           `json:"view"`                     // 播放量
	Like       int64           `json:"like"`                     // 喜欢
	Comment    int64           `json:"comment"`                  // 评论数
	Favorite   int64           `json:"favorite"`                 // 收藏
	Duration   int64           `json:"duration"`                 // 时长 /秒
	Hot        int64           `json:"hot"`
	VideoCount int64           `json:"video_count"`
	UserId     int64           `json:"user_id"`
	UserName   string          `json:"user_name"`
	UserAvatar string          `json:"user_avatar"`
	State      CollectionState `json:"state"`
	PublishAt  *timex.Time     `json:"publish_at"`
	CreatedAt  timex.Time      `json:"created_at"`
	UpdatedAt  timex.Time      `json:"updated_at"`
	Orders     string          `json:"orders"`

	User   *User    `json:"user"`
	Videos []*Video `json:"videos"`
}

type CollectionType uint8

const (
	CollectionTypeDiv      CollectionType = 1 // 自制
	CollectionTypeTransfer CollectionType = 2 // 转载
	CollectionTypeAnime    CollectionType = 3 // 动漫
	CollectionTypeSeries   CollectionType = 4 // 连续剧
)

type CollectionState uint8

const (
	CollectionStateNormal   CollectionState = 1 // 正常展示
	CollectionStateHidden   CollectionState = 2 // 隐藏
	CollectionStateUserBan  CollectionState = 3 // 禁止
	CollectionStateAdminBan CollectionState = 4 // admin禁止
)

type Video struct {
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	Cover        string      `json:"cover"` // 封面
	Brief        string      `json:"brief"` // 简介
	Url          string      `json:"url"`
	View         int64       `json:"view"`     // 播放量
	Like         int64       `json:"like"`     // 喜欢
	Comment      int64       `json:"comment"`  // 评论数
	Favorite     int64       `json:"favorite"` // 收藏
	Duration     int64       `json:"duration"` // 时长 /秒
	CollectionId int64       `json:"collection_id"`
	CreatedAt    timex.Time  `json:"created_at"`
	UpdatedAt    timex.Time  `json:"updated_at"`
	ShowTime     *timex.Time `json:"publish_at"`

	Collection *Collection `json:"collection"`
}
