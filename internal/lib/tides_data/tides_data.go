package tides_data

import (
	"bytes"
	"log"
	"time"

	tides_data_client "github.com/troptropcontent/what_the_tide/internal/lib/tides_data/client"
	tides_data_parser "github.com/troptropcontent/what_the_tide/internal/lib/tides_data/parser"
	"golang.org/x/net/html"
)

type ExtractFromWeb struct {
	Date   time.Time
	PortId int
	Tides  []tides_data_parser.Tide
}

func (extract *ExtractFromWeb) Load() {
	var html_bytes []byte
	tides_data_client.LoadWebPage(extract.Date, extract.PortId, &html_bytes)

	html_document, err := html.Parse(bytes.NewReader(html_bytes))
	if err != nil {
		log.Fatal(err)
	}

	tides_data_parser.ExtractTidesFromHtml(html_document, &extract.Tides)
}
