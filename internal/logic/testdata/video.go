package testdata

import (
	"net/http"
	"video_web/internal/consts"
)

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

var VideoTests = []struct {
	In       *videoIn
	Expected *videoExpected
}{
	{
		In:       &videoIn{Name: "三只松鼠", Type: consts.VideoTypeNormal, Category: 1, Cover: "", Brief: "", Episodes: []string{"www.abc.com", "www.cdf.com", "www.333.com"}, DelIndex: 2},
		Expected: &videoExpected{AddCode: http.StatusOK, DelCode: http.StatusOK},
	},
}
