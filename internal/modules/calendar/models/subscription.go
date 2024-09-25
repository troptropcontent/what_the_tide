package calendar_models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	CalendarID uint
	Email      string
	Published  bool
	ExternalId *string
}

func NewSubscription(calendarID uint, email string) *Subscription {
	return &Subscription{
		CalendarID: calendarID,
		Email:      email,
		Published:  false,
	}
}
