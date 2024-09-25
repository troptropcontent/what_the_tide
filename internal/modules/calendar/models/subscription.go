package calendar_models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	AgendaId   uint
	Email      string
	Published  bool
	ExternalId *string
}

func NewSubscription(agendaId uint, email string) *Subscription {
	return &Subscription{
		AgendaId:  agendaId,
		Email:     email,
		Published: false,
	}
}
