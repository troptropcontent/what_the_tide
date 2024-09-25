package google_calendar

import (
	"fmt"

	"google.golang.org/api/calendar/v3"
)

func CalendarList(service *calendar.Service) (*calendar.CalendarList, error) {
	request := service.CalendarList.List()

	_, err := service.Calendars.Get("toto").Do()

	if err != nil {
		fmt.Errorf("error retrieving something list : %v", err)
	}

	list, err := request.Do()
	if err != nil {
		fmt.Errorf("error retrieving calendar list : %v", err)
	}
	return list, nil
}

func CalendarCreate(service *calendar.Service, name string, description string) (*calendar.Calendar, error) {
	calendar := calendar.Calendar{
		Summary:     name,
		Description: description,
	}
	request := service.Calendars.Insert(&calendar)

	return request.Do()
}

func CalendarShareWith(service *calendar.Service, calendarId string, email string) (aclRule *calendar.AclRule, err error) {
	return service.Acl.Insert(calendarId, &calendar.AclRule{
		Scope: &calendar.AclRuleScope{
			Type:  "user",
			Value: email,
		},
		Role: "reader",
	}).Do()
}
