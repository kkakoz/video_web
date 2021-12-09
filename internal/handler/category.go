package handler

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/internal/dto"
	"google.golang.org/grpc/metadata"
)

type CategoryHandler struct {
	categoryLogic domain.ICategoryLogic
}

func NewCategoryHandler(categoryLogic domain.ICategoryLogic) *CategoryHandler {
	return &CategoryHandler{categoryLogic: categoryLogic}
}

func (item CategoryHandler) Add(ctx echo.Context) error {
	md := metadata.New(nil)
	newCtx := metadata.NewIncomingContext(ctx.Request().Context(), md)
	req := &dto.CategoryAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	category := &domain.Category{}
	err = copier.Copy(category, req)
	if err != nil {
		return err
	}
	err = item.categoryLogic.Add(newCtx, category)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item CategoryHandler) List(ctx echo.Context) error {
	md := metadata.New(nil)
	newCtx := metadata.NewIncomingContext(ctx.Request().Context(), md)
	list, err := item.categoryLogic.List(newCtx)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}