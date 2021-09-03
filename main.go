package main

import (
	"flag"
	"fmt"
	"go-sitemap/link"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*
	TODO
	- get the webpage
	- parse links
	- build urls with links
	- filter links with diff domain
	- find all pages(using bfs)
	- print xml
*/

func main() {
	urlFlag := flag.String("url", "https://blog.r-ush.co", "url to see the sitemap for")

	flag.Parse()
	fmt.Println("searching for ---> ", *urlFlag)

	pages := getPage(*urlFlag)

	for _, page := range pages {
		fmt.Println(page)
	}
}

func getPage(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return filterLinks(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)

	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		default:
			fmt.Println("skipping this--> ", l)
		}
	}

	return ret
}

func filterLinks(links []string, keepFn func(string) bool) []string {
	var ret []string

	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
