package scraper

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Anchor represents
type Anchor struct {
	href string
}

func getHTML(url string) string {
	response, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(bytes)
}

// RenderNode ...
func RenderNode(node *html.Node) string {
	var buf bytes.Buffer
	writer := io.Writer(&buf)
	html.Render(writer, node)
	return buf.String()
}

func crawler(node *html.Node, tag string, tags *[]Anchor) {
	if node.Type == html.ElementNode {
		if node.Data == tag {
			// To-Do: Ajeitar isso (ta feio)
			nodeString := RenderNode(node)
			hrefIndex := strings.Index(nodeString, "href")
			hrefBegin := hrefIndex + 6
			hrefEnd := strings.Index(nodeString[hrefBegin:], "\"")
			href := nodeString[hrefBegin : hrefBegin+hrefEnd]
			anchor := Anchor{href: href}
			(*tags) = append(*tags, anchor)
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		crawler(child, tag, tags)
	}

}

// GetAnchors ...
func GetAnchors(url string) ([]Anchor, error) {
	htm := getHTML(url)
	doc, err := html.Parse(strings.NewReader(htm))
	if err != nil {
		return nil, errors.New("Error parsing HTML")
	}
	var tags []Anchor
	crawler(doc, "a", &tags)
	return tags, nil
}
