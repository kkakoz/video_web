package router

import (
	"encoding/json"
	"github.com/kkakoz/pkg/redisx"
	"github.com/labstack/echo"
	"time"
	"video_web/internal/model/entity"
	"video_web/internal/pkg/keys"
	"video_web/internal/pkg/local"
)

func authority(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		user, err := local.GetUser(ctx.Request().Context())
		if err != nil {
			return err
		}

		// 日活跃用户
		redisx.Client().SetBit(ctx.Request().Context(), keys.DailyActiveUserKey(time.Now()), user.ID, 1)
		return f(ctx)
	}
}

// 设置登录用户
func setUser(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		if token == "" {
			return f(ctx)
		}
		client := redisx.Client()
		user := &entity.User{}
		result, err := client.Get(ctx.Request().Context(), keys.TokenKey(token)).Result()
		if err != nil {
			return nil
		}
		err = json.Unmarshal([]byte(result), user)
		if err != nil {
			return nil
		}
		//ctx.Request().Header.Add(local.UserLocalKey, result)
		ctx.SetRequest(ctx.Request().WithContext(local.WithUserLocal(ctx.Request().Context(), user)))
		return f(ctx)
	}
}

// UniqueVisitor 独立访客,ip统计数据
func UniqueVisitor(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ip := ctx.RealIP()
		redisx.Client().PFAdd(ctx.Request().Context(), keys.UniqueVisitorKey(time.Now()), ip)
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
