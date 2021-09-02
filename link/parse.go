package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a tag of the HTML doc
type Link struct {
	Href string
	Text string
}

// Parses the HTML document to slice of links
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)

	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	// dfs(doc, "")

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = "TODO:parse"
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

// func dfs(n *html.Node, padding string) {
// 	fmt.Println(padding, n.Data)

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding+"  ")
// 	}
// }

// testing
var exampleHtml = `
<html>
<body>
<h1>Hello!</h1>
<a href="/other-page">A link to another page
<span> some span  </span>
</a>
<a href="/page-two">A link to a second page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
