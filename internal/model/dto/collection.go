package dto

import (
	"video_web/internal/model/entity"
	"video_web/pkg/timex"
)

type VideoAdd struct {
	Name       string           `json:"name" validate:"required"`
	Type       entity.VideoType `json:"type" `
	CategoryId int64            `json:"category_id" validate:"required"`
	Cover      string           `json:"cover" validate:"required"` // 封面
	Brief      string           `json:"brief" validate:"required"`
	PublishAt  *timex.Time      `json:"publish_at" `
	//Videos     []VideoEasy           `json:"videos"`
}

//type VideoEasy struct {
//	Url      string `json:"url"`
//	Name     string `json:"name"`
//	Duration int64  `json:"duration"`
//}

type VideoId struct {
	VideoId int64 `json:"video_id" query:"video_id"`
}

type BackResourceList struct {
	CategoryId uint   `json:"category_id"`
	OrderType  uint8  `json:"order_type"` // 0默认时间排序 1热度排序
	Name       string `json:"name"`
	Pager
}

type BackCollectionList struct {
	OrderType uint8  `json:"order_type" query:"order_type"` // 0默认时间排序 1热度排序
	Name      string `json:"name" query:"name"`
	Pager
}
