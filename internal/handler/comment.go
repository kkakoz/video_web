package handler

import (
	"github.com/labstack/echo"
	"video_web/internal/domain"
	"video_web/internal/dto/request"
	"video_web/internal/pkg/mdctx"
)

type CommentHandler struct {
	commentLogic domain.ICommentLogic
}

func NewCommentHandler(commentLogic domain.ICommentLogic) *CommentHandler {
	return &CommentHandler{commentLogic: commentLogic}
}

func (item *CommentHandler) Add(ctx echo.Context) error {
	req := &request.CommentAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = item.commentLogic.Add(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return nil
}

func (item *CommentHandler) AddSubComment(ctx echo.Context) error {
	req := &request.SubCommentAddReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = item.commentLogic.AddSub(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return nil
}

func (item *CommentHandler) Get(ctx echo.Context) error {
	req := &request.CommentListReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := item.commentLogic.GetList(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}

func (item *CommentHandler) GetSubComment(ctx echo.Context) error {
	req := &request.SubCommentListReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	list, err := item.commentLogic.GetSubList(mdctx.NewCtx(ctx.Request()), req)
	if err != nil {
		return err
	}
	return ctx.JSON(200, list)
}
