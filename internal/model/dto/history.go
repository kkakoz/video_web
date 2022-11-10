package dto

type HistoryAdd struct {
	VideoId    int64  `json:"video_id"`
	ResourceId int64  `json:"resource_id"`
	Duration   int64  `json:"duration"`
	IP         string `json:"ip"`
}

type HistoryList struct {
	LastId int64 `query:"last_id"`
}
