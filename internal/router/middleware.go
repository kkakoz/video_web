package router

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github/kkakoz/video_web/internal/domain"
	"github/kkakoz/video_web/internal/pkg/keys"
	"github/kkakoz/video_web/pkg/errno"
	"github/kkakoz/video_web/pkg/redis"
)

func authority(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("X-Authorization")
		if token == "" {
			return errno.NewErr(401, 401, "请重新登录")
		}
		client := redis.GetClient()
		user := &domain.User{}
		result, err := client.Get(keys.TokenKey(token)).Result()
		if err != nil {
			return errno.NewErr(401, 401, "请重新登录")
		}
		err = json.Unmarshal([]byte(result), user)
		if err != nil {
			return errno.NewErr(401, 401, "请重新登录")
		}
		ctx.Set("user", user)
		return f(ctx)
	}
}

func setAccessOriginUrl(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return f(ctx)
	}
}
