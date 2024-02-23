package api

import (
	"context"
	"io"
	"os"
	"path/filepath"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/juancwu/jellyfan-web/views/component"
	"github.com/labstack/echo/v4"
)

func Upload(c echo.Context) error {
	// name := c.FormValue(component.UPLOAD_FORM_FILE_NAME)

	file, err := c.FormFile(component.UPLOAD_FORM_FILE_BLOB)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	src, err := file.Open()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	defer src.Close()

	// var custonName string
	// if name != "" {
	// 	custonName = name
	// } else {
	// 	custonName = file.Filename
	// }

	// destPath := filepath.Join(os.Getenv("UPLOAD_DIR"), custonName)
	destPath := filepath.Join(os.Getenv("UPLOAD_DIR"), file.Filename)
	dest, err := os.Create(destPath)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	alertId, err := gonanoid.New(6)
	if err != nil {
		return err
	}

	c.Response().Header().Set("HX-Reswap", "afterbegin")
	c.Response().Header().Set("HX-Retarget", "#alerts")

	return component.Alert("Successfully uploaded file!", "success", alertId).Render(context.Background(), c.Response().Writer)
}
