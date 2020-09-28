package scraper

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// Anchor represents
type Anchor struct {
	Href string
}

func getHTML(url string) string {
	response, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(bytes)
}

func renderNode(node *html.Node) string {
	var buf bytes.Buffer
	writer := io.Writer(&buf)
	html.Render(writer, node)
	return buf.String()
}

func getHref(nodeString string) string {
	hrefIndex := strings.Index(nodeString, "href")
	hrefBegin := hrefIndex + 6
	hrefEnd := strings.Index(nodeString[hrefBegin:], "\"")
	return nodeString[hrefBegin : hrefBegin+hrefEnd]
}

func crawler(node *html.Node, tag string, tags *[]Anchor) {
	if node.Type == html.ElementNode {
		if node.Data == tag {
			// To-Do: Ajeitar isso (ta feio)
			nodeString := renderNode(node)
			href := getHref(nodeString)
			if matched, _ := regexp.MatchString("^/wiki/.*", href); matched {
				anchor := Anchor{
					Href: href,
				}
				(*tags) = append(*tags, anchor)
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		crawler(child, tag, tags)
	}

}

// GetTags ...
func GetTags(url string, tag string) ([]Anchor, error) {
	htm := getHTML(url)
	doc, err := html.Parse(strings.NewReader(htm))
	if err != nil {
		return nil, errors.New("Error parsing HTML")
	}
	var tags []Anchor
	crawler(doc, tag, &tags)
	return tags, nil
}
