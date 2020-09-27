package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getBody(html string) string {
	bodyIndexStart := strings.Index(html, "<body")
	bodyIndexEnd := strings.Index(html, "</body>")
	return html[bodyIndexStart : bodyIndexEnd+7]
}

func getHTML(url string) string {
	response, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(bytes)
}

func main() {
	url := flag.String("url", "https://google.com/", "Url to begin parsing")
	flag.Parse()

	fmt.Printf("Url: %s\n", *url)
	html := getHTML(*url)
	body := getBody(html)
	print(body)
}
