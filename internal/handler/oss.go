package handler

import (
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type OssHandler struct {
}

func NewOssHandler() *OssHandler {
	return &OssHandler{}
}

func (item *OssHandler) GetConf(ctx echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"region":          viper.Get("oss.region"),
		"accessKeyId":     viper.Get("oss.accessKeyId"),
		"accessKeySecret": viper.Get("oss.accessKeySecret"),
		"bucket":          viper.Get("oss.bucket"),
	})
}
