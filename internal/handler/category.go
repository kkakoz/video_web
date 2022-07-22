package handler

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/dto/request"
	"video_web/internal/logic"
	"video_web/internal/model"
)

type categoryHandler struct {
}

var categoryOnce sync.Once
var _category *categoryHandler

func Category() *categoryHandler {
	categoryOnce.Do(func() {
		_category = &categoryHandler{}
	})
	return _category
}

func (item categoryHandler) Add(ctx echo.Context) error {
	req := &request.CategoryAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	category := &model.Category{}
	err = copier.Copy(category, req)
	if err != nil {
		return err
	}
	err = logic.Category().Add(ctx.Request().Context(), category)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item categoryHandler) List(ctx echo.Context) error {
	list, err := logic.Category().List(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}
