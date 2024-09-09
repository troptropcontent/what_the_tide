package database

import (
	"github.com/troptropcontent/what_the_tide/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DBPath = "database/database.db"

func MustInit() {
	db, err := gorm.Open(sqlite.Open(config.Root()+"/"+DBPath), &gorm.Config{})
	if err != nil {
		panic("error while connecting to database")
	}
	DB = db
}
