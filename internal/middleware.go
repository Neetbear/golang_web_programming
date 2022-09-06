package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func BasicLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[method=${method}, uri=${uri}, status=${status}]\n",
	})
}

func BodyLogger() echo.MiddlewareFunc {
	return middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Fprintf(os.Stdout, "[requestBody: %s]\n[responseBody: %s]\n", reqBody, resBody)
	})
}
