package router

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"github.com/kkakoz/pkg/logger"
	"github.com/kkakoz/video-rpc/pkg/errs"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strings"
	"video_web/pkg/errno"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func ErrHandler() echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {

		tar := &errno.Err{}
		if errors.As(err, &tar) {
			ctx.JSON(tar.HttpCode, tar)
			return
		}

		if httpErr := new(echo.HTTPError); errors.As(err, &httpErr) {
			ctx.JSON(httpErr.Code, map[string]any{
				"code": httpErr.Code,
				"msg":  httpErr.Message,
			})
			return
		}
		//if se := new(errs.Error); errors.As(err, &se) {
		//	ctx.JSON(400, map[string]any{
		//		"code": se.Code,
		//		"msg":  se.Message,
		//	})
		//}

		if _, ok := status.FromError(err); ok {
			newErr := errs.FromError(err)
			ctx.JSON(int(newErr.Code), map[string]any{
				"code": newErr.Code,
				"msg":  newErr.Message,
			})
			return
		}

		if validatorErrs, ok := err.(validator.ValidationErrors); ok {
			errs := []string{}
			for _, fieldErr := range validatorErrs {
				errs = append(errs, fieldErr.Translate(translator))
			}
			ctx.JSON(400, map[string]interface{}{
				"code": 400,
				"msg":  strings.Join(errs, ","),
			})
			return
		}

		fmt.Printf("%+v\n", err)
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

type customValidator struct {
	validate *validator.Validate
}

var translator ut.Translator

func init() {
	e := en.New()
	uniTrans := ut.New(e, e, zh.New(), zh_Hant_TW.New())
	translator, _ = uniTrans.GetTranslator("zh")
}

func NewValidator() customValidator {
	validate := validator.New()
	zh2.RegisterDefaultTranslations(validate, translator)
	return customValidator{validate: validate}
}

func (v customValidator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}
