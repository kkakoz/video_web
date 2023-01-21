package handler

import (
	"github.com/kkakoz/pkg/echox"
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
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
	return echox.OK(ctx)
}

func (collectHandler) List(ctx echo.Context) error {
	req := &dto.CollectList{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := logic.Collect().List(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"data": list,
	})
}
