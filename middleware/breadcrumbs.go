package middleware

import "github.com/labstack/echo/v4"

func ParseBreadcrumbs(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		unparsedBreadcrumbs := c.QueryParam("crumbs")

	}
}
