package httptest__test

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http/httptest"
)

type Res struct {
	*httptest.ResponseRecorder
}

func testGet(target string, ) Res {
	req := httptest.NewRequest("GET", target, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	if userToken != "" {
		req.Header.Set("X-Authorization", userToken)
	}
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)
	return Res{res}
}

func testPost(target string, body interface{}) Res {
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", target, bytes.NewReader(bodyBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	if userToken != "" {
		req.Header.Set("X-Authorization", userToken)
	}
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)
	return Res{res}
}

func testPut(target string, body interface{}) Res {
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest("PUT", target, bytes.NewReader(bodyBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	if userToken != "" {
		req.Header.Set("X-Authorization", userToken)
	}
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)
	return Res{res}
}

func testDel(target string) Res {
	req := httptest.NewRequest("DELETE", target, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	if userToken != "" {
		req.Header.Set("X-Authorization", userToken)
	}
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)
	return Res{res}
}

