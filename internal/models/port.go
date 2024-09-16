package models

import (
	"gorm.io/gorm"
)

const PortsConfigFile = "config/ports.json"

type Port struct {
	gorm.Model
	ExternalId int
	Name       string
}
