package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

// Link is a link
type Link struct {
	Href string
	Text string
}

// Links are links
var Links []Link

func main() {
	fileName := flag.String("file", "example1.html", "file to parse")
	flag.Parse()
	content := readFile(*fileName)
	r := strings.NewReader(content)

	document, e := html.Parse(r)
	if e != nil {
		panic(e)
	}
	var links []Link
	parseNode(document, &links)

	for _, link := range links {
		fmt.Printf("Href: %+v\n", link.Href)
		fmt.Printf("Text: %+v\n", link.Text)
		fmt.Println("  ")
	}

	for _, l := range Links {
		fmt.Printf("%+v\n", l.Href)
		fmt.Printf("%+v\n", l.Text)
		fmt.Println(" ")
	}
}

func readFile(file string) string {
	bytes, e := ioutil.ReadFile(file)
	if e != nil {
		panic(e)
	}

	return string(bytes)
}

func parseNode(node *html.Node, links *[]Link) string {
	var text string

	if node != nil {
		if node.Data == "a" {
			parseATag(node, links)
		}
		if node.FirstChild == nil {

			if node.Type == html.TextNode {
				trimmed := strings.TrimSpace(node.Data)
				if len(trimmed) > 0 {
					text = text + trimmed
				}
			}
		}

		if node.NextSibling != nil {
			data := parseNode(node.NextSibling, links)
			text = text + data
		}

		if node.FirstChild != nil {
			data := parseNode(node.FirstChild, links)
			text = text + data
		}

	}
	return text
}

func parseATag(node *html.Node, links *[]Link) {
	link := Link{}
	link.Href = getHref(node.Attr)
	link.Text = parseNode(node.FirstChild, links)
	*links = append(*links, link)
}

func getHref(attrs []html.Attribute) string {
	var s string
	for _, attr := range attrs {
		if attr.Key == "href" {
			s = attr.Val
		}
	}
	return s
}
