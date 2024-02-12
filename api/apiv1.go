package api

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/juancwu/jellyfan-web/views/page"
	"github.com/labstack/echo/v4"
)

const (
    UploadFormFileName = "filename"
    UploadFormCategory = "category"
    UploadFormFile = "file"
)

func Upload(c echo.Context) error {
    name := c.FormValue(UploadFormFileName)
    category := c.FormValue(UploadFormCategory)

    file, err := c.FormFile(UploadFormFile)
    if err != nil {
        return err
    }

    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    var custonName string
    if name != "" {
        custonName = name
    } else {
        custonName = file.Filename
    }

    destPath := filepath.Join(os.Getenv("UPLOAD_DIR"), fmt.Sprintf("%s.%s", custonName, category))
    dest, err := os.Create(destPath)
    if err != nil {
        return err
    }
    defer dest.Close()

    _, err = io.Copy(dest, src)
    if err != nil {
        return err
    }

    page.SuccessPage().Render(context.Background(), c.Response().Writer)
    return nil
}
