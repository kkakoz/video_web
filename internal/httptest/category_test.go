package httptest__test

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/tidwall/gjson"
	"github/kkakoz/video_web/internal/dto"
	"github/kkakoz/video_web/internal/httptest/testdata"
	"net/http"
	"testing"
)

func TestCategory(t *testing.T) {
	convey.Convey("category", t, func() {
		convey.Convey("add category", func() {
			for _, tt := range testdata.CategoryTests {
				req := &dto.CategoryAddReq{Name: tt.In}
				res := testPost("/categories", req)
				convey.So(res.Code, convey.ShouldEqual, tt.Expected)
			}
		})

		convey.Convey("category list", func() {
			res := testGet("/categories")
			convey.So(res.Code, convey.ShouldEqual, http.StatusOK)
			resstr := fmt.Sprint(res.Body)
			for i, tt := range testdata.CategoryTests {
				name := gjson.Get(resstr, fmt.Sprintf("%d.name", i)).String()
				convey.So(name, convey.ShouldEqual, tt.In)
			}

		})
	})
}
