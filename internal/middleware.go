package internal

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human}\n`,
	})
}

func BodyLogger() echo.MiddlewareFunc {
	return middleware.BodyDump(func(ctx echo.Context, b1, b2 []byte) {
		fmt.Fprintf(os.Stdout, "${%s} | ${%s}", b1, b2)
	})
}
