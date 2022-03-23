package echox_test

import (
	"fmt"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http/httptest"
	"testing"
	"video_web/pkg/echox"
)

func TestEcho(t *testing.T) {
	server := echo.New()
	server.Binder = echox.NewBinder()
	server.Validator = echox.NewValidator()
	// server.HTTPErrorHandler = echox.ErrHandler()
	server.Logger = echox.NewLogger(zap.Logger{})
	server.GET("/user/:id", func(context echo.Context) error {
		type Req struct {
			Id int `uri:"id" validate:"gte=15,lte=130"`
		}
		req := &Req{}
		err := context.Bind(req)
		if err != nil {
			// if e, ok := err.(validator.ValidationErrors); ok {
			//	for _, fieldErr := range e {
			//		fmt.Println(fieldErr.Error())
			//	}
			// }
			return err
		}
		return nil
	})
	request := httptest.NewRequest("GET", "/user/11", nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, request)
	fmt.Println(recorder.Body)
}
