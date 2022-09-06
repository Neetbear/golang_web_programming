package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

/*
Get /memberships/:id

Get /memberships

Post /memberships

Post /memberships/:id/points

*/

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
