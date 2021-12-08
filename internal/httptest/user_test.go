package httptest__test

import (
	"github.com/smartystreets/goconvey/convey"
	"github/kkakoz/video_web/internal/dto"
	"github/kkakoz/video_web/internal/httptest/testdata"
	"net/http"
	"testing"
)

func TestUser(t *testing.T) {
	convey.Convey("Test User CRUD", t, func() {
		for _, data := range testdata.UserTestData {
			convey.Convey("Test User Register", func() {
				req := &dto.RegisterReq{
					Name:         data.Req.Name,
					IdentityType: data.Req.IdentifyType,
					Identifier:   data.Req.Identifier,
					Credential:   data.Req.Credential,
				}
				res := testPost("/users", req)
				convey.So(res.Code, convey.ShouldEqual, http.StatusOK)
			})

			convey.Convey("Test User Login", func() {
				req := &dto.LoginReq{
					IdentityType: data.Req.IdentifyType,
					Identifier:   data.Req.Identifier,
					Credential:   data.Req.Credential,
				}
				res := testPost("/users/login", req)
				convey.So(res.Code, convey.ShouldEqual, http.StatusOK)
			})
		}

	})

}
