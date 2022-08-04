package dto

type HistoryAdd struct {
	VideoId  int64  `json:"video_id"`
	Duration int64  `json:"duration"`
	IP       string `json:"ip"`
}
