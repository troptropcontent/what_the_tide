package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func healthCheckHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func main() {
	e := echo.New()

	e.File("/", "internal/views/index.html")
	e.GET("/up", healthCheckHandler)
	e.Static("/public", "public")
	e.Logger.Fatal(e.Start(":3001"))
}
