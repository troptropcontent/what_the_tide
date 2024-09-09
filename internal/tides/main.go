package tides

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Tide struct {
	Time  time.Time
	High  bool
	Level int
}

func Get(date time.Time) (tides []Tide, err error) {
	parsedTides := map[string]Tide{}
	var correctReportDate *bool

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		r.Headers.Set("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
		r.Headers.Set("Cache-Control", "no-cache")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Cookie", "PHPSESSID=ecd7h5ecjfod2m5vqht5ao7fe1; UserAgreement=7fd4a8ff68d527f145d812c13b5ac262dd9acabf2af62adca72635236905766ef2caad66; __utma=1.414734129.1725606961.1725606961.1725606961.1; __utmc=1; __utmz=1.1725606961.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmv=1.|2=ads=1=1; __eoi=ID=962d4d590f4c5344:T=1725606959:RT=1725662146:S=AA-Afjayt49_938q-iwx5OEZ7rZU")
		r.Headers.Set("DNT", "1")
		r.Headers.Set("Pragma", "no-cache")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Sec-Fetch-Site", "none")
		r.Headers.Set("Sec-Fetch-User", "?1")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("sec-ch-ua", `"Not;A=Brand";v="24", "Chromium";v="128"`)
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", `"Windows"`)
	})

	c.OnHTML("#MareeEnteteJour", func(e *colly.HTMLElement) {
		expectedReportDate := fmt.Sprintf("%d %s %d", date.Day(), Months[date.Month()-1], date.Year())
		r := strings.Contains(e.DOM.Text(), expectedReportDate)
		correctReportDate = &r
	})

	c.OnHTML("td.SEPV>.HauteursInfoFlag", func(e *colly.HTMLElement) {
		timeAttr, _ := e.DOM.Attr("data-hm")
		timeRegext := regexp.MustCompile(`^(?P<hours>\d{2})h(?P<minutes>\d{2})_(?P<type>(t|b))$`)
		timeMatches := timeRegext.FindStringSubmatch(timeAttr)
		if len(timeMatches) == 0 {
			//   Handle errors
		}
		hoursString := timeMatches[1]
		hoursInt, _ := strconv.Atoi(hoursString)
		minutesString := timeMatches[2]
		minutesInt, _ := strconv.Atoi(minutesString)
		tideTime := time.Date(date.Year(), date.Month(), date.Day(), hoursInt, minutesInt, 0, 0, date.Location())
		high := timeMatches[3] == "b"

		levelAttr, _ := e.DOM.Attr("data-ht")
		levelRegext := regexp.MustCompile(`(\d{1,2}),(\d{1,2})m`)
		levelMatches := levelRegext.FindStringSubmatch(levelAttr)
		if len(levelMatches) == 0 {
			//   Handle errors
		}

		levelMeters := levelMatches[1]
		levelCentimeters := levelMatches[2]
		levelInCentimeters, _ := strconv.Atoi(levelMeters + levelCentimeters)

		tide := Tide{
			Time:  tideTime,
			High:  high,
			Level: levelInCentimeters,
		}

		_, tideAlreadyRegistered := parsedTides[hoursString+minutesString]

		if !tideAlreadyRegistered {
			parsedTides[hoursString+minutesString] = tide
			tides = append(tides, tide)
		}
	})

	base_url := os.Getenv("TIDE_WEBSITE_BASE_URL")
	url := fmt.Sprintf("%v/%d?d=%s", base_url, Ports.Biarritz, date.Format("20060102"))

	c.Visit(url)

	if *correctReportDate {
		return tides, nil
	}

	return []Tide{}, err
}
