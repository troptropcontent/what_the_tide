package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/troptropcontent/what_the_tide/config"
	"github.com/troptropcontent/what_the_tide/database"
	"github.com/troptropcontent/what_the_tide/internal/lib/tides_data"
	"github.com/troptropcontent/what_the_tide/internal/models"
)

func main() {
	loc, _ := time.LoadLocation("Europe/Paris")
	defaultDateFlag := time.Now().In(loc).AddDate(0, 0, 7).Format("20060102")
	datePtr := flag.String("date", defaultDateFlag, "the date of the extraction, must be provided in the following format YYYYMMDD")
	flag.Parse()

	date, err := time.ParseInLocation("20060102", *datePtr, loc)
	if err != nil {
		fmt.Println(err)
		panic("date flag should be in the following format YYYYMMDD")
	}

	// Connect to database
	database.MustInit()

	// Run migrations related to models
	models.MustInit()

	// Load the list of the ports for wich we want to extract the data
	portConfig := models.PortsConfig{}
	portConfig.LoadFromJson(config.Root() + "/" + models.PortsConfigFile)

	// Save each new tides to the database
	for _, port := range portConfig.Ports {
		tides := []models.Tide{}
		tides_data.Get(date, port.Id, &tides)
		database.DB.Create(tides)
	}
}
