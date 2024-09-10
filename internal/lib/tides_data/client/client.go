package tides_data_client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func LoadWebPage(date time.Time, port_id int, html_file *[]byte) {
	base_url := os.Getenv("TIDE_WEBSITE_BASE_URL")
	url := fmt.Sprintf("%v/%d?d=%s", base_url, port_id, date.Format("20060102"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Adding headers
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "PHPSESSID=ecd7h5ecjfod2m5vqht5ao7fe1; UserAgreement=7fd4a8ff68d527f145d812c13b5ac262dd9acabf2af62adca72635236905766ef2caad66; __utma=1.414734129.1725606961.1725606961.1725606961.1; __utmc=1; __utmz=1.1725606961.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmv=1.|2=ads=1=1; __eoi=ID=962d4d590f4c5344:T=1725606959:RT=1725662146:S=AA-Afjayt49_938q-iwx5OEZ7rZU")
	req.Header.Add("DNT", "1")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", `"Not;A=Brand";v="24", "Chromium";v="128"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"Windows"`)

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
