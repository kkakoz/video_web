package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
)

type commentHandler struct {
}

var commentOnce sync.Once
var _comment *commentHandler

func Comment() *commentHandler {
	commentOnce.Do(func() {
		_comment = &commentHandler{}
	})
	return _comment
}

func (item *commentHandler) Add(ctx echo.Context) error {
	req := &dto.CommentAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Comment().Add(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return nil
}

func (item *commentHandler) AddSubComment(ctx echo.Context) error {
	req := &dto.SubCommentAdd{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Comment().AddSub(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return nil
}

func (item *commentHandler) Get(ctx echo.Context) error {
	req := &dto.CommentList{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := logic.Comment().GetList(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *commentHandler) GetSubComment(ctx echo.Context) error {
	req := &dto.SubCommentList{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := logic.Comment().GetSubList(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *commentHandler) Delete(ctx echo.Context) error {
	req := &dto.CommentDel{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Comment().Delete(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *commentHandler) DeleteSubComment(ctx echo.Context) error {
	req := &dto.SubCommentDel{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = logic.Comment().DeleteSubComment(ctx.Request().Context(), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}
