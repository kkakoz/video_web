package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
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

func (historyHandler) AddHistory(ctx echo.Context) error {
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
	return nil
}

func (historyHandler) List(ctx echo.Context) error {
	return nil
}
