package handler

import (
	"github.com/kkakoz/pkg/echox"
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
)

type noticeHandler struct {
}

var noticeOnce sync.Once
var _notice *noticeHandler

func Notice() *noticeHandler {
	noticeOnce.Do(func() {
		_notice = &noticeHandler{}
	})
	return _notice
}

func (noticeHandler) UnReadCount(ctx echo.Context) error {
	count, err := logic.Notice().UnReadCount(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"count": count,
	})
}

func (noticeHandler) UpdateNotice(ctx echo.Context) error {
	req := &dto.ID{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Notice().ReadNotice(ctx.Request().Context(), req.ID)
	if err != nil {
		return err
	}
	return echox.Ok(ctx)
}

func (noticeHandler) GetList(ctx echo.Context) error {
	req := &dto.LastId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	list, err := logic.Notice().GetList(ctx.Request().Context(), req.LastId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"data": list,
	})
}

func (noticeHandler) ReadAll(ctx echo.Context) error {
	err := logic.Notice().ReadAll(ctx.Request().Context())
	if err != nil {
		return err
	}

	return echox.OK(ctx)
}
