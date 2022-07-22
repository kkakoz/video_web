package handler

import (
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"sync"
)

type ossHandler struct {
}

var ossOnce sync.Once
var _oss *ossHandler

func Oss() *ossHandler {
	ossOnce.Do(func() {
		_oss = &ossHandler{}
	})
	return _oss
}

func (item *ossHandler) GetConf(ctx echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"region":          viper.Get("oss.region"),
		"accessKeyId":     viper.Get("oss.accessKeyId"),
		"accessKeySecret": viper.Get("oss.accessKeySecret"),
		"bucket":          viper.Get("oss.bucket"),
	})
}
