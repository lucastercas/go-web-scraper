package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	args := os.Args[1:]
	url := args[0]

	fmt.Printf("Url: %s\n", url)
	html := getHTML(url)
	body := getBody(html)
	print(body)
}
