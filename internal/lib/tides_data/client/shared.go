package tides_data_client

import (
	"log"
	"net/http"
)

func BuildRequestWithHeaders(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("error during request instanciation: ", err)
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

	return req
}
