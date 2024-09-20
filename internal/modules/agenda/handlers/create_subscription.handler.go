package agenda_handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/what_the_tide/internal/lib/worker"
)

func CreateSubscription(c echo.Context) error {
	worker.Push("SubscribeToAlreadyExistingAgenda", []string{})
	return c.File("internal/views/subscription/created.html")
}
