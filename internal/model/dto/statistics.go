package dto

import (
	"video_web/pkg/timex"
)

type CalculateUV struct {
	StartAt timex.Time `json:"start_at"`
	EndAt   timex.Time `json:"end_at"`
}

type CalculateDAU struct {
	StartAt timex.Time `json:"start_at"`
	EndAt   timex.Time `json:"end_at"`
}
