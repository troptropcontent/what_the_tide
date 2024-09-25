package calendar_models

import (
	"testing"

	"github.com/troptropcontent/what_the_tide/database"
)

func TestNewBasicAgenda(t *testing.T) {
	t.Run("It assigns the correct type to the record", func(t *testing.T) {
		if got := NewBasicCalendar(); got.Type != BasicCalendarType {
			t.Errorf("NewBasicAgenda() should set the type to %v, got %v", BasicCalendarType, got.Type)
		}
	})
}

func TestAssociations(t *testing.T) {
	t.Run("Configuration", func(t *testing.T) {
		database.MustInit()
		transaction := database.DB.Begin()
		t.Cleanup(func() { transaction.Rollback() })
		var calendars []BasicCalendar
		numberOfCalendarsBefore := transaction.Find(&calendars).RowsAffected
		var calendarConfigurations []BasicCalendarConfiguration
		numberOfConfigurationsBefore := transaction.Find(&calendarConfigurations).RowsAffected
		calendar := NewBasicCalendar()
		calendar.Configuration = BasicCalendarConfiguration{
			PortID: 123,
		}
		transaction.Create(&calendar)
		numberOfCalendarsAfter := transaction.Find(&calendars).RowsAffected
		numberOfConfigurationsAfter := transaction.Find(&calendarConfigurations).RowsAffected
		t.Run("It should be possible to create a configuration through the calendar", func(t *testing.T) {
			if numberOfCalendarsBefore+1 != numberOfCalendarsAfter {
				numberOfCalendarsCreated := numberOfCalendarsAfter - numberOfCalendarsBefore
				t.Errorf("expected exactly 1 calendar created got : %d", numberOfCalendarsCreated)
			}
			if numberOfConfigurationsBefore+1 != numberOfConfigurationsAfter {
				numberOfConfigurationsCreated := numberOfConfigurationsAfter - numberOfCalendarsBefore
				t.Errorf("expected exactly 1 configuration created got : %d", numberOfConfigurationsCreated)
			}
		})
		t.Run("It should be possible to retrieve the configuration through the calendar", func(t *testing.T) {
			transaction.Joins("Configuration").Find(&calendar, calendar.ID)
			if calendar.Configuration.PortID != 123 {
				t.Errorf("The Configuration field have not been loaded")
			}
		})
	})
}
