package echox

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github/kkakoz/video_web/pkg/errno"
)

func ErrHandler() echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		if validatorErr, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(200, validatorErr.Translate(translator))
			return
		}
		tar := &errno.Err{}
		if errors.As(err, tar) {
			ctx.JSON(tar.HttpCode, tar)
			return
		}
		ctx.JSON(500, err.Error())
	}
}
