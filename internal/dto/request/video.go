package request

import "video_web/pkg/timex"

type VideoAddReq struct {
	Name       string        `json:"name" validate:"required"`
	Type       uint8         `json:"type" validate:"required"`
	CategoryId int64         `json:"category_id" validate:"required"`
	Cover      string        `json:"cover" validate:"required"` // 封面
	Brief      string        `json:"brief" validate:"required"`
	PublishAt  *timex.Time   `json:"publish_at" validate:"required"`
	EpisodeIds []int64       `json:"episode_ids"`
	Episodes   []EpisodeEasy `json:"episodes"`
}

type EpisodeEasy struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type EpisodeAddReq struct {
	Url        string `json:"url"`
	Name       string `json:"name"`
	Cover      string `json:"cover"` // 封面
	VideoId    int64  `uri:"video_id"`
	CategoryId int64  `json:"category_id"`
	Brief      string `json:"brief"`
	AddType    uint8  `json:"add_type"`
}

type VideoIdReq struct {
	VideoId int64 `uri:"video_id"`
}

type EpisodeIdReq struct {
	VideoId   int64 `uri:"video_id"`
	EpisodeId int64 `uri:"episode_id"`
}

type VideosReq struct {
	CategoryId uint  `query:"category_id"`
	LastValue  uint  `query:"last_value"`
	OrderType  uint8 `query:"order_type"`
}

type BackVideosReq struct {
	CategoryId uint  `query:"category_id"`
	OrderType  uint8 `query:"order_type"`
	Pager
}
