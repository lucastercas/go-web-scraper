package racer

import (
	"fmt"
	"regexp"

	scraper "github.com/lucastercas/web-scraper/internal/scraper"
)

func race(wiki string, curArticle string, finalArticle string, curDepth int, finalDepth int) bool {
	if curArticle == finalArticle {
		return true
	}
	if curDepth == finalDepth {
		return false
	}
	return race(wiki, curArticle, finalArticle, curDepth+1, finalDepth)
}

// Race ...
func Race(wiki string, firstArticle string, finalArticle string, finalDepth int) {
	fmt.Printf("Racing from on [%s]: [%s] => [%s]\n", wiki, firstArticle, finalArticle)

	anchors, _ := scraper.GetTags(wiki+"/wiki/"+firstArticle, "a")
	curDepth := 0
	for _, anchor := range anchors {
		re := regexp.MustCompile(`/wiki/(?P<url>.*)$`)
		res := re.FindStringSubmatch(anchor.Href)
		curArticle := res[1]
		if result := race(wiki, curArticle, finalArticle, curDepth+1, finalDepth); result {
			fmt.Printf("%s == %s", curArticle, finalArticle)
		}
	}
}
