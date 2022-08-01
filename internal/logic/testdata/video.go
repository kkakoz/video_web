package testdata

type videoIn struct {
	Name         string   `json:"name"`
	Type         int32    `json:"type"`
	Category     int32    `json:"category"` // 分类
	Cover        string   `json:"cover"`    // 封面
	Brief        string   `json:"brief"`
	Episodes     []string `json:"episode"`
	DelIndex     int      `json:"del_index"`      // 删除的index
	DelEpisodeId int64    `json:"delete_episode"` // 删除的id
}

type videoExpected struct {
	AddCode int `json:"add_code"`
	DelCode int `json:"del_code"`
}
