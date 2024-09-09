package main

import (
	"fmt"
	"log"
	"time"

	"github.com/troptropcontent/what_the_tide/config"
	"github.com/troptropcontent/what_the_tide/database"
	"github.com/troptropcontent/what_the_tide/internal/lib/google_calendar"
	"github.com/troptropcontent/what_the_tide/internal/models"
	"google.golang.org/api/calendar/v3"
)

func publishTideEvent(tide *models.Tide, srv *calendar.Service, portsConfig *models.PortsConfig) error {
	loc, _ := time.LoadLocation("Europe/Paris")
	port, err := portsConfig.FindPort(tide.PortId)
	if err != nil {
		log.Fatalf("Unable to retrieve port data: %v", err)
	}

	calendarId, err := google_calendar.FetchCalendarId(srv, port.Name)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar id: %v", err)
	}

	baseName := "LOW"
	baseDescription := "Lowest"
	if tide.High {
		baseName = "HIGH"
		baseDescription = "Highest"
	}
	eventName := fmt.Sprintf("%s - %s", baseName, tide.Time.In(loc).Format(time.Kitchen))
	eventDescription := fmt.Sprintf("%s point : %.2f m", baseDescription, float32(tide.Level)/float32(100.0))
	event := &calendar.Event{
		Summary:     eventName,
		Location:    port.Name,
		Description: eventDescription,
		Start: &calendar.EventDateTime{
			DateTime: tide.Time.Add(-10 * time.Minute).Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: tide.Time.Add(10 * time.Minute).Format(time.RFC3339),
		},
	}

	err = google_calendar.CreateEvent(srv, calendarId, event)
	if err != nil {
		log.Fatalf("Unable to create event : %v", err)
	}
	tide.IsPublished = true
	database.DB.Save(&tide)
	return nil
}

func main() {
	// Connect to database
	database.MustInit()

	//  Load google calendar service
	calendarService, err := google_calendar.LoadService()
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	// Load Ports data
	portConfig := models.PortsConfig{}
	portConfig.LoadFromJson(config.Root() + "/" + models.PortsConfigFile)

	// Load all unpublished tides
	unpublishedTides := []models.Tide{}
	database.DB.Where(&models.Tide{IsPublished: false}).Find(&unpublishedTides)

	for _, unpublishedTide := range unpublishedTides {
		publishTideEvent(&unpublishedTide, calendarService, &portConfig)
	}
}
