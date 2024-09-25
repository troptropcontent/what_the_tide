package calendar_jobs

import (
	"context"
	"fmt"

	"github.com/troptropcontent/what_the_tide/database"
	"github.com/troptropcontent/what_the_tide/internal/lib/google_calendar"
	calendar_models "github.com/troptropcontent/what_the_tide/internal/modules/calendar/models"
)

type SubscribeToAlreadyExistingAgendaJob struct{}

func (j *SubscribeToAlreadyExistingAgendaJob) Name() string {
	return "SubscribeToAlreadyExistingAgendaJob"
}

func (j *SubscribeToAlreadyExistingAgendaJob) Perform(ctx context.Context, args ...interface{}) error {
	subscriptionId := args[0].(string)
	subscription := calendar_models.Subscription{}
	database.DB.Limit(1).Find(&subscription, subscriptionId)
	if subscription.ID == 0 {
		return fmt.Errorf("the subscription with id %v does not exists", subscriptionId)
	}

	agenda := calendar_models.Base{}
	database.DB.Limit(1).Find(&agenda, subscription.AgendaId)
	if subscription.ID == 0 {
		return fmt.Errorf("the agenda with id %v does not exists", subscription.AgendaId)
	}

	service, err := google_calendar.NewService()
	if err != nil {
		return fmt.Errorf("error building calendar service: %v", err)
	}

	acl, err := google_calendar.CalendarShareWith(service, agenda.ExternalId, subscription.Email)
	if err != nil {
		return fmt.Errorf("error sharing calendar: %v", err)
	}

	subscription.ExternalId = &acl.Id
	database.DB.Save(&subscription)

	return nil
}
