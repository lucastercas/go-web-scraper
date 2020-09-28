package main

import (
	"flag"
	"fmt"
	"os"

	scraper "github.com/lucastercas/web-scraper/internal/scraper"
)

func main() {
	url := flag.String("url", "https://pt.wikipedia.org/wiki/Hunting", "Url to begin parsing")
	flag.Parse()

	fmt.Printf("Url: %s\n", *url)
	tags, err := scraper.GetAnchors(*url)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(tags)

}
