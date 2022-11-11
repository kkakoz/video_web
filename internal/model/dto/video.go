package dto

type ResourceAdd struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	Brief    string `json:"brief"`
	Duration int64  `json:"duration"`
}

type ResourceAddList struct {
	VideoId   int64         `json:"video_id" `
	Resources []ResourceAdd `json:"resources"`
}

type ResourceId struct {
	ResourceId int64 `json:"resource_id"`
}

type Videos struct {
	CategoryId int64  `query:"category_id"`
	LastId     int64  `query:"last_id"`
	OrderType  uint8  `query:"order_type"`
	UserId     int64  `query:"user_id"`
	Search     string `query:"search"`
}
