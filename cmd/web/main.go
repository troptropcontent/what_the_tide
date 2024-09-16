package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func healthCheckHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func createSubscriptionHandler(c echo.Context) error {
	return c.File("internal/views/subscription/created.html")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.File("/", "internal/views/index.html")
	e.GET("/up", healthCheckHandler)
	e.Static("/public", "public")

	e.POST("/subscription", createSubscriptionHandler)
	
	e.Logger.Fatal(e.Start(":3001"))
}
