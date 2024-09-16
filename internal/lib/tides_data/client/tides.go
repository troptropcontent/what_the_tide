package tides_data_client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func LoadPortWebPage(date time.Time, port_id int, html_file *[]byte) {
	base_url := os.Getenv("WHAT_THE_TIDE_TIDE_WEBSITE_BASE_URL")
	if base_url == "" {
		log.Fatalf("WHAT_THE_TIDE_TIDE_WEBSITE_BASE_URL seems to be not set and is required")
	}

	url := fmt.Sprintf("%v/%d?d=%s", base_url, port_id, date.Format("20060102"))
	req := BuildRequestWithHeaders(url)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	*html_file = body
}
