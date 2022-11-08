package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/pkg/echox"
)

type collectHandler struct {
}

var collectOnce sync.Once
var _collect *collectHandler

func Collect() *collectHandler {
	collectOnce.Do(func() {
		_collect = &collectHandler{}
	})
	return _collect
}

func (collectHandler) Add(ctx echo.Context) error {
	req := &dto.CollectAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Collect().Add(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return echox.Ok(ctx)
}
