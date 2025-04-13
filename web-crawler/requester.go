package main

import (
	"io"
	"log"
	"net/http"
)

const MaxReadLength = 1.5e+7

func readResponseBody(r *http.Response) string {
	limitedReader := io.LimitReader(r.Body, MaxReadLength)
	data, err := io.ReadAll(limitedReader)
	if err != nil {
		log.Fatalln(err)
	}

	return string(data)

}

func getUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		return ""
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing http stream, %v", err)
		}
	}(resp.Body)

	return readResponseBody(resp)
}

func crawlerGetUrl(url string, URLList *URLList) {
	response := getUrl(url)
	URLList.AddToURLListFromHtml(response)
}
