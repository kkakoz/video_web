package echox

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func ErrHandler() echo.HTTPErrorHandler {
	return func(err error, context echo.Context) {
		if validatorErr, ok := err.(validator.ValidationErrors); ok {
			context.JSON(200, validatorErr.Translate(translator))
		}
	}
}
