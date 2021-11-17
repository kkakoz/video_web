package echox

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

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
