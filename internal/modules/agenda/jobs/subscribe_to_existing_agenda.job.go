package agenda_jobs

import (
	"context"
	"fmt"

	"github.com/troptropcontent/what_the_tide/internal/lib/google_calendar"
)

func SubscribeToAlreadyExistingAgenda(ctx context.Context, calendarId string, email string) error {
	service, err := google_calendar.NewService()
	if err != nil {
		return fmt.Errorf("error building calendar service: %v", err)
	}

	err = google_calendar.CalendarShareWith(service, calendarId, email)
	if err != nil {
		return fmt.Errorf("error sharing calendar: %v", err)
	}
	return nil
}
