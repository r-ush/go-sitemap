package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://blog.r-ush.co", "url to see the sitemap for")

	flag.Parse()
	fmt.Println(*urlFlag)

	/*
		TODO
		- get the webpage
		- parse links
		- build urls with links
		- filter links with diff domain
		- find all pages(using bfs)
		- print xml
	*/
}
