package dto

type ResourceAdd struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	Brief    string `json:"brief"`
	Duration int64  `json:"duration"`
	Show     bool   `json:"show"`
}

type ResourceId struct {
	ResourceId int64 `json:"resource_id"`
}

type Videos struct {
	CategoryId uint  `json:"category_id"`
	LastValue  uint  `json:"last_value"`
	OrderType  uint8 `json:"order_type"`
}
