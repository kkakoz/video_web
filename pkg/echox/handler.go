package echox

import (
	"fmt"
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
		fmt.Printf("%+v\n", err)
		ctx.JSON(500, map[string]interface{}{
			"code": 500,
			"msg":  err.Error(),
		})
	}
}
