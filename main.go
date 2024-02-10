package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

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
		page.LandingPage().Render(context.Background(), c.Response().Writer)
		return nil
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
