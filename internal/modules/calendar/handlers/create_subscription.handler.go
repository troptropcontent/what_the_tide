package calendar_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/troptropcontent/what_the_tide/database"
	calendar_models "github.com/troptropcontent/what_the_tide/internal/modules/calendar/models"
)

type CreateSubscriptionParams struct {
	PortId uint   `form:"port_id"`
	Email  string `form:"email"`
}

func CreateSubscription(c echo.Context) error {
	var params CreateSubscriptionParams
	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	var calendar calendar_models.BasicAgenda
	database.DB.Joins("Configuration").Where("configurations.port_id = ?", params.PortId).Limit(1).Find(&calendar)
	// TODO add validation with go validator here
	// Does a BasicAgenda exists already for this Port ?
	// Yes => We just need to add a new ACL to the email to this agenda
	// No => We need to :
	// - Create one (BasicAgenda record + BasicAgendaConfiguraion)
	// - add a new ACL to the email to this agenda
	// - trigger the publication of tides
	return c.File("internal/views/subscription/created.html")
}
