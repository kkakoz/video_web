package dto

type DanmuAdd struct {
	Content    string `json:"content"`
	ResourceId string `json:"resource_id"`
	Duration   int64  `json:"duration"`
	Color      string `json:"color"`
}
