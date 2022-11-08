package echox

import "github.com/labstack/echo"

func Ok(ctx echo.Context) error {
	return ctx.JSON(200, map[string]any{
		"code": 200,
		"msg":  "ok",
	})
}
