package handler

import (
	"github.com/kkakoz/pkg/echox"
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
)

type danmuHandler struct {
}

var danmuOnce sync.Once
var _danmu *danmuHandler

func Danmu() *danmuHandler {
	danmuOnce.Do(func() {
		_danmu = &danmuHandler{}
	})
	return _danmu
}

func (danmuHandler) Send(ctx echo.Context) error {
	req := &dto.DanmuAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Danmu().Add(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return echox.OK(ctx)
}
