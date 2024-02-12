package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/juancwu/jellyfan-web/api"
	"github.com/juancwu/jellyfan-web/route"
	"github.com/juancwu/jellyfan-web/views/component"
	"github.com/juancwu/jellyfan-web/views/page"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Jellyfan Web")

	e := echo.New()

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		page.LandingPage([]component.Crumb{}).Render(context.Background(), c.Response().Writer)
		return nil
	})

    apiGroup := e.Group(route.API_V1_PREFIX)
    apiGroup.POST(route.API_V1_UPLOAD_SUBPATH, api.Upload)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
