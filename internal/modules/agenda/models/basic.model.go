package agenda_models

import (
	"gorm.io/gorm"
)

const BasicAgendaType = "basic"

type BasicAgendaConfiguration struct {
	gorm.Model
	PortID int
}

type BasicAgenda struct {
	Base
	Configuration BasicAgendaConfiguration
}

func NewBasicAgenda() *BasicAgenda {
	return &BasicAgenda{Base: Base{Type: BasicAgendaType}}
}
