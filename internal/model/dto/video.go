package dto

type VideoAdd struct {
	Url        string `json:"url"`
	Name       string `json:"name"`
	Cover      string `json:"cover"` // 封面
	CategoryId int64  `json:"category_id"`
	Brief      string `json:"brief"`
	Duration   int64  `json:"duration"`
	Show       bool   `json:"show"`
}

type VideoId struct {
	VideoId int64 `query:"video_id"`
}

type Videos struct {
	CategoryId uint  `query:"category_id"`
	LastValue  uint  `query:"last_value"`
	OrderType  uint8 `query:"order_type"`
}

type BackVideoList struct {
	CategoryId uint   `json:"category_id"`
	OrderType  uint8  `json:"order_type"` // 0默认时间排序 1热度排序
	Name       string `json:"name"`
	Pager
}
