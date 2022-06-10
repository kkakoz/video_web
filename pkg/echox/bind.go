package echox

import (
	"errors"
	"github.com/labstack/echo"
	"reflect"
	"strconv"
)

type binder struct {
	binder echo.Binder
}

func NewBinder() binder {
	return binder{binder: &echo.DefaultBinder{}}
}

func (b binder) Bind(v interface{}, c echo.Context) error {
	err := b.binder.Bind(v, c)
	if err != nil {
		return err
	}
	err = bindUri(v, c)
	if err != nil {
		return err
	}
	err = c.Validate(v)
	return err
}

func bindUri(ptr interface{}, c echo.Context) error {
	typ := reflect.TypeOf(ptr).Elem()
	val := reflect.ValueOf(ptr).Elem()
	for i := 0; i < typ.NumField(); i++ {
		tfield := typ.Field(i)
		uri := tfield.Tag.Get("uri")
		if uri == "" {
			continue
		}
		param := c.Param(uri)
		if param == "" {
			continue
		}
		vfield := val.Field(i)
		switch vfield.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			atoi, err := strconv.Atoi(param)
			if err != nil {
				return errors.New("bind field " + vfield.String() + "failed")
			}
			vfield.SetInt(int64(atoi))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			atoi, err := strconv.Atoi(param)
			if err != nil {
				return errors.New("bind field " + vfield.String() + "failed")
			}
			vfield.SetUint(uint64(atoi))
		case reflect.String:
			vfield.SetString(param)
		default:
			return errors.New("bind failed")
		}
	}
	return nil
}
