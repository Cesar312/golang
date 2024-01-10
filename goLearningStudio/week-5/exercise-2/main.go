package main

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
)

func ScrapeWiki(url string) (stext string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodys := string(body)
	t_start := strings.Index(bodys, `<p><b>Socrates</b> (<span class="rt-commentedText nowrap">`)
	t_end := strings.Index(bodys[t_start:], "</p>")

	body2 := bodys[t_start : t_start+t_end]

	return body2
}

func main() {
	const url = "https://en.wikipedia.org/wiki/Socrates"
	myScrape := ScrapeWiki(url)
	fmt.Println(html.UnescapeString(myScrape))
}
