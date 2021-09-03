package main

import (
	"flag"
	"fmt"
	"go-sitemap/link"
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

	resp, err := http.Get(*urlFlag)
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

	links, _ := link.Parse(resp.Body)

	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		default:
			fmt.Println("skipping this--> ", l)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)

	}

}
