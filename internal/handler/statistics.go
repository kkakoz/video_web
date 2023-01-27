package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
)

type statisticsHandler struct {
}

var statisticsOnce sync.Once
var _statistics *statisticsHandler

func Statistics() *statisticsHandler {
	statisticsOnce.Do(func() {
		_statistics = &statisticsHandler{}
	})
	return _statistics
}

func (statistics *statisticsHandler) CalculateUV(ctx echo.Context) error {

	req := &dto.CalculateUV{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	res, err := logic.Statistics().CalculateUV(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, res)
}

func (statistics *statisticsHandler) CalculateDAU(ctx echo.Context) error {

	req := &dto.CalculateDAU{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	res, err := logic.Statistics().CalculateDAU(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, res)
}
