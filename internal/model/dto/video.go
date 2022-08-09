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
	VideoId int64 `json:"video_id"`
}

type Videos struct {
	CategoryId uint  `json:"category_id"`
	LastValue  uint  `json:"last_value"`
	OrderType  uint8 `json:"order_type"`
}
