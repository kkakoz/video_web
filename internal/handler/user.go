package handler

import (
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/internal/dto"
	"google.golang.org/grpc/metadata"
)

type UserHandler struct {
	userLogic domain.IUserLogic
}

func NewUserHandler(userLogic domain.IUserLogic) *UserHandler {
	return &UserHandler{userLogic: userLogic}
}

func (item *UserHandler) Login(ctx echo.Context) error {
	md := metadata.New(nil)
	newCtx := metadata.NewIncomingContext(ctx.Request().Context(), md)
	auth := &dto.LoginReq{}
	err := ctx.Bind(auth)
	if err != nil {
		return err
	}
	token, err := item.userLogic.Login(newCtx, auth)
	if err != nil {
		return err
	}
	return ctx.JSON(200, token)
}

func (item *UserHandler) Register(ctx echo.Context) error {
	md := metadata.New(nil)
	newCtx := metadata.NewIncomingContext(ctx.Request().Context(), md)
	auth := &dto.RegisterReq{}
	err := ctx.Bind(auth)
	if err != nil {
		return err
	}
	err = item.userLogic.Register(newCtx, auth)
	if err != nil {
		return err
	}
	return ctx.JSON(200, nil)
}

func (item *UserHandler) GetUser(ctx echo.Context) error {
	return ctx.JSON(200, nil)
}
