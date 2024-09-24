package google_calendar

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/troptropcontent/what_the_tide/config"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func credentialsPath() string {
	return config.Root() + "/credentials/service_account_credentials.json"
}

func Client() (client *http.Client, err error) {
	ctx := context.Background()
	jsonCredentials, err := os.ReadFile(credentialsPath())
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(jsonCredentials, calendar.CalendarScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	return config.Client(ctx), nil
}
