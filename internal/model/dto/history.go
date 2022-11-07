package dto

type HistoryAdd struct {
	ResourceId int64  `json:"resource_id"`
	Duration   int64  `json:"duration"`
	IP         string `json:"ip"`
}
