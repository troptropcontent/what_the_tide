package models

import (
	"time"

	"gorm.io/gorm"
)

type Tide struct {
	gorm.Model
	PortId      int
	Time        time.Time
	High        bool
	Level       int
	IsPublished bool
}
