package agenda

import "github.com/labstack/echo/v4"

func CreateSubscriptionHandler(c echo.Context) error {
	return c.File("internal/views/subscription/created.html")
}
