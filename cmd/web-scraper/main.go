package main

import (
	"flag"

	"github.com/lucastercas/web-scraper/internal/racer"
)

func main() {
	wiki := flag.String("wiki", "https://pt.wikipedia.org", "Wiki to race")
	begin := flag.String("begin", "Franceses", "Article to begin racing")
	end := flag.String("end", "Galeses", "Article to end racing")
	depth := flag.Int("depth", 3, "Depth to go for finding")
	flag.Parse()

	racer.Race(*wiki, *begin, *end, *depth)

}

// Franceses -> Galeses
// Franceses -> Galeses -> Roma_Antiga
