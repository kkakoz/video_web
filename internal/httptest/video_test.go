package httptest__test

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/tidwall/gjson"
	"github/kkakoz/video_web/internal/dto"
	"github/kkakoz/video_web/internal/httptest/testdata"
	"testing"
)

func TestVideo(t *testing.T) {
	convey.Convey("video crud", t, func() {
		convey.Convey("add video", func() {
			loginUser()
			defer logout()
			for _, tt := range testdata.VideoTests {
				req := &dto.AddVideoReq{
					Name:     tt.In.Name,
					Type:     tt.In.Type,
					Category: tt.In.Category,
					Cover:    tt.In.Cover,
					Brief:    tt.In.Brief,
				}
				res := testPost("/videos", req)
				convey.So(res.Code, convey.ShouldEqual, tt.Expected.AddCode)
			}
		})

		convey.Convey("video add episode", func() {
			i := 1
			loginUser()
			defer logout()
			for _, tt := range testdata.VideoTests {
				if tt.Expected.AddCode == 200 {
					for _, episode := range tt.In.Episodes {
						req := &dto.AddEpisodeReq{
							Url: episode,
						}
						res := testPost(fmt.Sprintf("/videos/%d/episodes", i), req)
						convey.So(res.Code, convey.ShouldEqual, 200)
					}
					i++
				}
			}
		})

		convey.Convey("get video", func() {
			videoId := 1
			loginUser()
			defer logout()
			for _, tt := range testdata.VideoTests {
				res := testGet(fmt.Sprintf("/videos/%d", videoId))
				convey.So(res.Code, convey.ShouldEqual, 200)
				resstr := fmt.Sprint(res.Body)
				convey.So(gjson.Get(resstr, "name").String(), convey.ShouldEqual, tt.In.Name)
				convey.So(gjson.Get(resstr, "episodes.0.pre_id").Int(), convey.ShouldEqual, 0)
				convey.So(gjson.Get(resstr, "episodes.1.pre_id").Int(), convey.ShouldEqual, 1)
				convey.So(gjson.Get(resstr, "episodes.0.next_id").Int(), convey.ShouldEqual, 2)
				convey.So(gjson.Get(resstr, "episodes.1.next_id").Int(), convey.ShouldEqual, 3)
				// 获取删除的id
				tt.In.DelEpisodeId = gjson.Get(resstr, fmt.Sprintf("episodes.%d.id", tt.In.DelIndex)).Int()
				videoId++
			}
		})

		convey.Convey("del episode", func() {
			loginUser()
			defer logout()
			for _, tt := range testdata.VideoTests {
				res := testDel(fmt.Sprintf("/episodes/%d", tt.In.DelEpisodeId))
				convey.So(res.Code, convey.ShouldEqual, tt.Expected.AddCode)
			}
		})
	})
}
