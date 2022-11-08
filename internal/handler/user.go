package handler

import (
	"github.com/labstack/echo"
	"sync"
	"video_web/internal/logic"
	"video_web/internal/model/dto"
	"video_web/internal/pkg/local"
	"video_web/pkg/echox"
)

type userHandler struct {
}

var userOnce sync.Once
var _user *userHandler

func User() *userHandler {
	userOnce.Do(func() {
		_user = &userHandler{}
	})
	return _user
}

func (item *userHandler) Login(ctx echo.Context) error {
	auth := &dto.Login{}
	err := ctx.Bind(auth)
	if err != nil {
		return err
	}
	token, err := logic.User().Login(ctx.Request().Context(), auth)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]any{
		"token": token,
	})
}

func (item *userHandler) Register(ctx echo.Context) error {
	auth := &dto.Register{}
	err := ctx.Bind(auth)
	if err != nil {
		return err
	}
	err = logic.User().Register(ctx.Request().Context(), auth)
	if err != nil {
		return err
	}
	return echox.Ok(ctx)
}

func (item *userHandler) GetCurUser(ctx echo.Context) error {
	user, err := local.GetUser(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(200, user)
}

func (item *userHandler) GetUser(ctx echo.Context) error {
	req := &dto.UserId{}
	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	user, err := logic.User().GetUser(ctx.Request().Context(), req.UserId)
	if err != nil {
		return err
	}
	return ctx.JSON(200, user)
}
