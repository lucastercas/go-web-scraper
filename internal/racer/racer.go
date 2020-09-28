package racer

import (
	"fmt"
	"regexp"

	scraper "github.com/lucastercas/web-scraper/internal/scraper"
)

func race(wiki string, begin string, end string, curDepth int, finalDepth int) bool {
	if begin == end {
		return true
	}
	return false
}

// Race ...
func Race(wiki string, begin string, end string, depth int) {
	fmt.Printf("Racing from on [%s]: [%s] => [%s]\n", wiki, begin, end)

	anchors, _ := scraper.GetTags(wiki+"/wiki/"+begin, "a")
	curDepth := 0
	for _, anchor := range anchors {
		re := regexp.MustCompile(`/wiki/(?P<url>.*)$`)
		res := re.FindStringSubmatch(anchor.Href)
		curArticle := res[1]
		if result := race(wiki, curArticle, end, curDepth+1, depth); result {
			fmt.Printf("%s == %s", curArticle, end)
		}
	}
}
