package model

type Danmu struct {
	ID      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	VideoId int64  `json:"video_id"`
	Mode    int    `json:"mode"`
	Text    string `json:"text"`
	Offset  int    `json:"offset"`
	Size    int    `json:"size"`
	Color   int64  `json:"color"`
}
