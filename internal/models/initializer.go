package models

import (
	"github.com/troptropcontent/what_the_tide/database"
)

func MustInit() {
	err := database.DB.AutoMigrate(
		&Tide{},
	)
	if err != nil {
		panic("error while migrating the database")
	}
}
