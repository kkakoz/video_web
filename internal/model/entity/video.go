package entity

import "video_web/pkg/timex"

type Video struct {
	ID            int64       `json:"id"`
	Name          string      `json:"name"`
	Type          VideoType   `json:"type"`
	UserId        int64       `json:"user_id" gorm:"index"`
	CategoryId    int64       `json:"category_id" gorm:"index"` // 分类
	Cover         string      `json:"cover"`                    // 封面
	Brief         string      `json:"brief"`                    // 简介
	View          int64       `json:"view"`                     // 播放量
	Like          int64       `json:"like"`                     // 喜欢
	Comment       int64       `json:"comment"`                  // 评论数
	Collect       int64       `json:"collect"`                  // 收藏
	Duration      int64       `json:"duration"`                 // 时长 /秒
	Hot           int64       `json:"hot"`
	ResourceCount int64       `json:"resource_count"`
	State         VideoState  `json:"state"`
	PublishAt     *timex.Time `json:"publish_at"`
	CreatedAt     timex.Time  `json:"created_at"`
	UpdatedAt     timex.Time  `json:"updated_at"`

	User      *User       `json:"user"`
	Resources []*Resource `json:"resources"`
	Category  Category    `json:"category"`
}

type VideoType uint8

const (
	VideoTypeDiv      VideoType = 1 // 自制
	VideoTypeTransfer VideoType = 2 // 转载
	VideoTypeAnime    VideoType = 3 // 动漫
	VideoTypeSeries   VideoType = 4 // 连续剧
)

type VideoState uint8

const (
	VideoStateDefault  VideoState = 0 // 未审核
	VideoStateNormal   VideoState = 1 // 正常展示
	VideoStateHidden   VideoState = 2 // 隐藏
	VideoStateUserBan  VideoState = 3 // 禁止
	VideoStateAdminBan VideoState = 4 // admin禁止
)

type Resource struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Brief     string     `json:"brief"` // 简介
	Url       string     `json:"url"`
	UserId    int64      `json:"user_id"`
	Duration  int64      `json:"duration"` // 时长 /秒
	VideoId   int64      `json:"video_id"`
	CreatedAt timex.Time `json:"created_at"`
	UpdatedAt timex.Time `json:"updated_at"`
	Sort      int64      `json:"sort"`
	Video     *Video     `json:"collection"`
	User      *User      `json:"user"`
}
