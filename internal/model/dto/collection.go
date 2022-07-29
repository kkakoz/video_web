package dto

import "video_web/pkg/timex"

type CollectionAdd struct {
	Name       string      `json:"name" validate:"required"`
	Type       uint8       `json:"type" `
	CategoryId int64       `json:"category_id" validate:"required"`
	Cover      string      `json:"cover" validate:"required"` // 封面
	Brief      string      `json:"brief" validate:"required"`
	PublishAt  *timex.Time `json:"publish_at" `
	Videos     []VideoEasy `json:"videos"`
}

type VideoEasy struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	Duration int64  `json:"duration"`
}

type CollectionId struct {
	CollectionId int64 `uri:"collection_id"`
}

type BackVideoList struct {
	CategoryId uint   `query:"category_id"`
	OrderType  uint8  `query:"order_type"` // 0默认时间排序 1热度排序
	Name       string `query:"name"`
	Pager
}

type BackCollectionList struct {
	OrderType uint8  `query:"order_type"` // 0默认时间排序 1热度排序
	Name      string `query:"name"`
	Pager
}
