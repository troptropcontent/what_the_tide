package database

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/troptropcontent/what_the_tide/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const BaseDbPath = "database/database"

func DbPath() string {
	var env string
	if env = os.Getenv("WHAT_THE_TIDE_ENV"); env == "" {
		log.Fatalf("environment variable WHAT_THE_TIDE_ENV seems not to be set")
	}

	return fmt.Sprintf("%s/%s_%s.db", config.Root(), BaseDbPath, env)
}

func MustInit() {
	db, err := gorm.Open(sqlite.Open(DbPath()), &gorm.Config{})
	if err != nil {
		panic("error while connecting to database")
	}
	DB = db
}
