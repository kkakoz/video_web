package echox

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
	"video_web/pkg/errno"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func ErrHandler(logger *zap.Logger) echo.HTTPErrorHandler {
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
		if errors.As(err, &tar) {
			ctx.JSON(tar.HttpCode, tar)
			return
		}
		logger.Error(err.Error(), zap.String("err", fmt.Sprintf("%+v", err)), zap.String("url", ctx.Request().RequestURI))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, map[string]interface{}{
				"code": 404,
				"msg":  err.Error(),
			})
			return
		}
		ctx.JSON(500, map[string]interface{}{
			"code": 500,
			"msg":  err.Error(),
		})
	}
}
