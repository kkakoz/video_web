package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/pkg/echox"
)

type historyHandler struct {
}

var historyOnce sync.Once
var _history *historyHandler

func History() *historyHandler {
	historyOnce.Do(func() {
		_history = &historyHandler{}
	})
	return _history
}

func (historyHandler) Add(ctx echo.Context) error {
	req := &dto.HistoryAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	req.IP = ctx.Request().RemoteAddr
	err = logic.History().Add(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return echox.OK(ctx)
}

func (historyHandler) List(ctx echo.Context) error {
	req := &dto.HistoryList{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := logic.History().List(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"data": list,
	})
}

func (historyHandler) Get(ctx echo.Context) error {
	req := &dto.VideoId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	history, err := logic.History().Get(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, history)
}
