package echox

import "github.com/labstack/echo"

func OK(ctx echo.Context) error {
	return ctx.JSON(200, map[string]any{
		"msg": "ok",
	})
}
