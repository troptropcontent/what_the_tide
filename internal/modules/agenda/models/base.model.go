package agenda_models

import "gorm.io/gorm"

type Base struct {
	gorm.Model
	Name        string
	ExternalId  string
	Type        string
	IsPublished bool
}

func (a *Base) TableName() string {
	return "agendas"
}
