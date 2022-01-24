package request

type AddVideoReq struct {
	Name     string `json:"name"`
	Type     int32  `json:"type"`
	Category int32  `json:"category"` // 分类
	Cover    string `json:"cover"`    // 封面
	Brief    string `json:"brief"`
}

type AddEpisodeReq struct {
	Url     string `json:"url"`
	VideoId int64  `uri:"video_id"`
	Order   int64  `json:"order"`
}

type VideoIdReq struct {
	VideoId int64 `uri:"video_id"`
}

type EpisodeIdReq struct {
	EpisodeId int64 `uri:"episode_id"`
}
