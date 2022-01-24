package handler

import (
	"github.com/labstack/echo"
	"video_web/internal/domain"
	"video_web/internal/dto/request"
	"video_web/internal/pkg/mdctx"
)

type UserHandler struct {
	userLogic domain.IUserLogic
}

func NewUserHandler(userLogic domain.IUserLogic) *UserHandler {
	return &UserHandler{userLogic: userLogic}
}

func (item *UserHandler) Login(ctx echo.Context) error {
	auth := &request.LoginReq{}
	err := ctx.Bind(auth)
	if err != nil {
		return err
	}
	token, err := item.userLogic.Login(mdctx.NewCtx(ctx.Request()), auth)
	if err != nil {
		return err
	}
	return ctx.JSON(200, token)
}

func (item *UserHandler) Register(ctx echo.Context) error {
	auth := &request.RegisterReq{}
	err := ctx.Bind(auth)
	if err != nil {
		return err
	}
	err = item.userLogic.Register(mdctx.NewCtx(ctx.Request()), auth)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *UserHandler) GetCurUser(ctx echo.Context) error {
	token := ctx.QueryParam("token")
	user, err := item.userLogic.GetCurUser(mdctx.NewCtx(ctx.Request()), token)
	if err != nil {
		return err
	}
	return ctx.JSON(200, user)
}

func (item *UserHandler) GetUser(ctx echo.Context) error {
	req := &request.UserReq{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	user, err := item.userLogic.GetUser(mdctx.NewCtx(ctx.Request()), req.UserId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, user)
}
