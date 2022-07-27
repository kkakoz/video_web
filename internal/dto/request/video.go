package request

import "video_web/pkg/timex"

type CollectionAddReq struct {
	Name       string        `json:"name" validate:"required"`
	Type       uint8         `json:"type" validate:"required"`
	CategoryId int64         `json:"category_id" validate:"required"`
	Cover      string        `json:"cover" validate:"required"` // 封面
	Brief      string        `json:"brief" validate:"required"`
	PublishAt  *timex.Time   `json:"publish_at" validate:"required"`
	Episodes   []EpisodeEasy `json:"episodes"`
}

type EpisodeEasy struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type VideoAddReq struct {
	Url        string `json:"url"`
	Name       string `json:"name"`
	Cover      string `json:"cover"` // 封面
	CategoryId int64  `json:"category_id"`
	Brief      string `json:"brief"`
	Duration   int64  `json:"duration"`
	Show       bool   `json:"show"`
}

type VideoIdReq struct {
	VideoId int64 `uri:"video_id"`
}

type VideosReq struct {
	CategoryId uint  `query:"category_id"`
	LastValue  uint  `query:"last_value"`
	OrderType  uint8 `query:"order_type"`
}

type BackVideosReq struct {
	CategoryId uint  `query:"category_id"`
	OrderType  uint8 `query:"order_type"` // 0默认时间排序 1热度排序
	Pager
}
