package httptest__test

import (
	"encoding/json"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/tidwall/gjson"
	"github/kkakoz/video_web/internal/dto"
	"github/kkakoz/video_web/internal/httptest/testdata"
	"net/http"
	"strconv"
	"testing"
)

func TestUser(t *testing.T) {
	convey.Convey("Test User CRUD", t, func() {

		convey.Convey("Test User Register", func() {
			for _, tt := range testdata.UserTests {
				req := &dto.RegisterReq{
					Name:         tt.In.Name,
					IdentityType: tt.In.IdentifyType,
					Identifier:   tt.In.Identifier,
					Credential:   tt.In.Credential,
				}
				res := testPost("/users", req)
				convey.So(res.Code, convey.ShouldEqual, http.StatusOK)
			}
		})

		convey.Convey("Test User Login", func() {
			for _, tt := range testdata.UserTests {
				req := &dto.LoginReq{
					IdentityType: tt.In.IdentifyType,
					Identifier:   tt.In.Identifier,
					Credential:   tt.In.Credential,
				}
				res := testPost("/users/login", req)
				convey.So(res.Code, convey.ShouldEqual, http.StatusOK)
				s := ""
				json.Unmarshal([]byte(fmt.Sprint(res.Body)), &s)
				tt.Expected.Token = s
			}
		})

		convey.Convey("Test Get Cur User", func() {
			for _, tt := range testdata.UserTests {
				// 在注册和登录成功的情况下测试当前用户
				if tt.Expected.RegisterCode == 200 && tt.Expected.LoginCode == 200 {
					res := testGet("/users/cur?token=" + tt.Expected.Token)
					resstr := fmt.Sprintf("%v", res.Body)
					convey.So(gjson.Get(resstr, "name").String(), convey.ShouldEqual, tt.In.Name)
				}
			}
		})

		convey.Convey("Test Get User", func() {
			i := 1
			loginUser()
			defer logout()
			for _, tt := range testdata.UserTests {
				// 在注册成功的情况下测试当前用户
				if tt.Expected.RegisterCode == 200 {
					// 用户登录
					res := testGet("/users/" + strconv.Itoa(i))
					resstr := fmt.Sprintf("%v", res.Body)
					convey.So(gjson.Get(resstr, "name").String(), convey.ShouldEqual, tt.In.Name)
					i++
				}
			}
		})
	})
}
