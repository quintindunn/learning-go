package main

import (
	"fmt"
	"log"
)

func main() {
	urls := []string{
		"https://wikipedia.com/wiki/web_crawler",
	}

	ulist, err := URLListFromArray(urls)
	if err != nil {
		log.Panicf("Error from generating URL list: %v\n", err)
	}

	for url := ulist.PopFront(); url != ""; url = ulist.PopFront() {
		crawlerGetUrl(url, ulist)
		fmt.Println(url)
	}
}
