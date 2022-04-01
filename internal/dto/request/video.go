package request

type VideoAddReq struct {
	Name       string        `json:"name"`
	Type       uint8         `json:"type"`
	CategoryId int64         `json:"category_id"`
	Cover      string        `json:"cover"` // 封面
	Brief      string        `json:"brief"`
	EpisodeIds []int64       `json:"episode_ids"`
	Episodes   []EpisodeEasy `json:"episodes"`
}

type EpisodeEasy struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type EpisodeAddReq struct {
	Url        string `json:"url"`
	Name       string `json:"name"`
	Cover      string `json:"cover"` // 封面
	VideoId    int64  `uri:"video_id"`
	CategoryId int64  `json:"category_id"`
	Brief      string `json:"brief"`
	AddType    uint8  `json:"add_type"`
}

type VideoIdReq struct {
	VideoId int64 `uri:"video_id"`
}

type EpisodeIdReq struct {
	VideoId   int64 `uri:"video_id"`
	EpisodeId int64 `uri:"episode_id"`
}

type VideosReq struct {
	CategoryId uint  `query:"category_id"`
	LastValue  uint  `query:"last_value"`
	OrderType  uint8 `query:"order_type"`
}
