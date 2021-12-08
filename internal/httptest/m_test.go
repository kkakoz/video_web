package httptest__test

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github/kkakoz/video_web/internal/handler"
	"github/kkakoz/video_web/internal/router"
	"github/kkakoz/video_web/pkg/conf"
	"github/kkakoz/video_web/pkg/mysqlx"
	"github/kkakoz/video_web/pkg/redis"
	"go.uber.org/fx"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server http.Handler

func TestMain(m *testing.M) {
	conf.InitTestConfig()
	if _, err := mysqlx.New(viper.GetViper()); err != nil {
		log.Fatalln("init mysql conn err:", err)
	}
	fx.New(
		handler.Provider,
		redis.Provider,
		router.Provider,
		fx.Supply(viper.GetViper()),
		fx.Populate(&server),
	)
	m.Run()
}

type Res struct {
	*httptest.ResponseRecorder
}

func testGet(target string, body interface{}) Res {
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest("GET", target, bytes.NewReader(bodyBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)
	return Res{res}
}

func testPost(target string, body interface{}) Res {
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", target, bytes.NewReader(bodyBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)
	return Res{res}
}

func testPut(target string, body interface{}) Res {
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest("PUT", target, bytes.NewReader(bodyBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)
	return Res{res}
}

func testDel(target string, body interface{}) Res {
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest("DELETE", target, bytes.NewReader(bodyBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)
	return Res{res}
}
