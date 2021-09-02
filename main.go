package main

import (
	"flag"
	"fmt"
	"net/http"
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
	// links, _ = link.parse()
}
