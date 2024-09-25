package calendar_models

import (
	"gorm.io/gorm"
)

const BasicCalendarType = "basic"

type BasicCalendarConfiguration struct {
	gorm.Model
	CalendarID uint
	PortID     int
}

type BasicCalendar struct {
	Base
	Configuration BasicCalendarConfiguration `gorm:"foreignKey:CalendarID"`
}

func NewBasicCalendar() *BasicCalendar {
	return &BasicCalendar{Base: Base{Type: BasicCalendarType}}
}
