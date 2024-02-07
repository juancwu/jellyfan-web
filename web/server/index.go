package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Serve() {
	e := echo.New()
	e.Static("/static", "web/static")

	e.GET("/", func(c echo.Context) error {
		c.Response().WriteHeader(http.StatusOK)
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
