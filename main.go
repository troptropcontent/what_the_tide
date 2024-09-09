package main

import (
	"fmt"
	"time"

	"github.com/troptropcontent/what_the_tide/internal/tides"
)

func main() {
	loc, _ := time.LoadLocation("Europe/Paris")
	date := time.Now().AddDate(0, 0, 7).In(loc)
	tides, err := tides.Get(date)
	if err == nil {
		for _, tide := range tides {
			highString := "Low"
			if tide.High {
				highString = "High"
			}
			fmt.Printf("%s - %s - %s\n", tide.Time.Format("2006 01 02"), tide.Time.Format(time.Kitchen), highString)
		}
	}
}
