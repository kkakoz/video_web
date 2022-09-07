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
	CategoryId int64 `json:"category_id"`
	LastValue  int64 `json:"last_value"`
	OrderType  uint8 `json:"order_type"`
	UserId     int64 `json:"user_id"`
}
