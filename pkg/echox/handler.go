package echox

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"strings"
	"video_web/pkg/errno"
)

func ErrHandler() echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		if validatorErrs, ok := err.(validator.ValidationErrors); ok {
			errs := []string{}
			for _, fieldErr := range validatorErrs {
				errs = append(errs, fieldErr.Translate(translator))
			}
			ctx.JSON(400, strings.Join(errs, ","))
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
