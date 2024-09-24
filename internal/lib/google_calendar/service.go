package google_calendar

import (
	"context"
	"fmt"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func NewService() (service *calendar.Service, err error) {
	ctx := context.Background()
	client, err := Client()
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}

	return calendar.NewService(ctx, option.WithHTTPClient(client))
}
