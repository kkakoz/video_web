package router

import (
	"encoding/json"
	"github.com/labstack/echo"
	"video_web/internal/model"
	"video_web/internal/pkg/keys"
	"video_web/internal/pkg/local"
	"video_web/pkg/errno"
	"video_web/pkg/redisx"
)

func authority(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		if token == "" {
			return errno.NewErr(401, 401, "请重新登录")
		}
		client := redisx.Client()
		user := &model.User{}
		result, err := client.Get(keys.TokenKey(token)).Result()
		if err != nil {
			return errno.NewErr(401, 401, "请重新登录")
		}
		err = json.Unmarshal([]byte(result), user)
		if err != nil {
			return errno.NewErr(401, 401, "请重新登录")
		}
		ctx.Request().Header.Add(local.UserLocalKey, result)
		ctx.SetRequest(ctx.Request().WithContext(local.NewCtx(ctx.Request())))
		return f(ctx)
	}
}

// func setAccessOriginUrl(f echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		middleware.CORS()
// 		method := ctx.Request().Method
// 		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
// 		ctx.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
// 		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
// 		ctx.Response().Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
// 		ctx.Response().Header().Set("Access-Control-Allow-Credentials", "true")
// 		if method == "OPTIONS" {
// 			return ctx.JSON(http.StatusNoContent, nil)
// 		}
// 		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
// 		return f(ctx)
// 	}
// }
