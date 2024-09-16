package tides_data

import (
	"bytes"
	"log"
	"time"

	tides_data_client "github.com/troptropcontent/what_the_tide/internal/lib/tides_data/client"
	ports_parser "github.com/troptropcontent/what_the_tide/internal/lib/tides_data/parser/ports"
	tides_parser "github.com/troptropcontent/what_the_tide/internal/lib/tides_data/parser/tides"
	"golang.org/x/net/html"
)

type TidesExtractFromWeb struct {
	Date   time.Time
	PortId int
	Tides  []tides_parser.Tide
}

func (extract *TidesExtractFromWeb) Load() {
	var html_bytes []byte
	tides_data_client.LoadPortWebPage(extract.Date, extract.PortId, &html_bytes)

	html_document, err := html.Parse(bytes.NewReader(html_bytes))
	if err != nil {
		log.Fatal(err)
	}

	tides_parser.ExtractTidesFromHtml(html_document, &extract.Tides)
}

type PortsExtractFromWeb struct {
	Ports []ports_parser.Port
}

func (extract *PortsExtractFromWeb) Load() {
	var html_bytes []byte
	tides_data_client.LoadPortsWebPage(&html_bytes)

	html_document, err := html.Parse(bytes.NewReader(html_bytes))
	if err != nil {
		log.Fatal(err)
	}

	ports_parser.ExtractPortsFromHtml(html_document, &extract.Ports)
}
