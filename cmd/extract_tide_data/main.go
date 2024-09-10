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
	tides := []models.Tide{}
	for _, port := range portConfig.Ports {
		extract := tides_data.ExtractFromWeb{
			Date:   date,
			PortId: port.Id,
		}

		extract.Load()

		for _, tide_from_web := range extract.Tides {
			tides = append(tides, models.Tide{
				Time:   extract.Date.AddDate(0, 0, tide_from_web.DaysOffset).Add(tide_from_web.Time).UTC(),
				PortId: port.Id,
				High:   tide_from_web.High,
				Level:  tide_from_web.Level,
				Coef:   tide_from_web.Coef,
			})
		}
	}

	database.DB.Save(&tides)
}
