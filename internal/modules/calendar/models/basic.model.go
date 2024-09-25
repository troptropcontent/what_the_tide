package calendar_models

import (
	"gorm.io/gorm"
)

const BasicAgendaType = "basic"

type BasicAgendaConfiguration struct {
	gorm.Model
	CalendarID uint
	PortID     int
}

type BasicAgenda struct {
	Base
	Configuration BasicAgendaConfiguration `gorm:"foreignKey:UserName"`
}

func NewBasicAgenda() *BasicAgenda {
	return &BasicAgenda{Base: Base{Type: BasicAgendaType}}
}
