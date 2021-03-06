package handler

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"video_web/internal/dto/request"
	"video_web/internal/logic"
	"video_web/internal/model"
	"video_web/internal/pkg/mdctx"
)

type CategoryHandler struct {
	categoryLogic *logic.CategoryLogic
}

func NewCategoryHandler(categoryLogic *logic.CategoryLogic) *CategoryHandler {
	return &CategoryHandler{categoryLogic: categoryLogic}
}

func (item CategoryHandler) Add(ctx echo.Context) error {
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
	err = item.categoryLogic.Add(mdctx.NewCtx(ctx.Request()), category)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item CategoryHandler) List(ctx echo.Context) error {
	list, err := item.categoryLogic.List(mdctx.NewCtx(ctx.Request()))
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}
