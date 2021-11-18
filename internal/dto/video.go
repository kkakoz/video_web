package dto

type AddVideoReq struct {
	Name      string `json:"name"`
	Type      int32  `json:"type"`
	Category  int32  `json:"category"` // 分类
	Cover     string `json:"cover"`    // 封面
	Brief     string `json:"brief"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
}

type AddEpisodeReq struct {
	Url     int64 `json:"url"`
	VideoId int64 `json:"video_id"`
}
