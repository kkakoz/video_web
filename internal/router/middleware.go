package router

import (
	"encoding/json"
	"github.com/labstack/echo"
	"video_web/internal/pkg/keys"
	"video_web/pkg/errno"
	"video_web/pkg/local"
	"video_web/pkg/redis"
)

func authority(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("X-Authorization")
		if token == "" {
			return errno.NewErr(401, 401, "请重新登录")
		}
		client := redis.GetClient()
		user := &local.User{}
		result, err := client.Get(keys.TokenKey(token)).Result()
		if err != nil {
			return errno.NewErr(401, 401, "请重新登录")
		}
		err = json.Unmarshal([]byte(result), user)
		if err != nil {
			return errno.NewErr(401, 401, "请重新登录")
		}
		ctx.Request().Header.Add(local.UserLocalKey, result)
		return f(ctx)
	}
}

func setAccessOriginUrl(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return f(ctx)
	}
}
