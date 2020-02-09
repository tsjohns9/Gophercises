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
}

func readFile(file string) string {
	bytes, e := ioutil.ReadFile(file)
	if e != nil {
		panic(e)
	}

	return string(bytes)
}

// Receives an html node and returns all the plain text it contains
// html nodes are trees. Recurse through all sub nodes to collect a slice of Link structs from the <a> tags
func parseNode(node *html.Node, links *[]Link) string {
	// all text collected from the original node passd into the function
	var text string

	if node != nil {
		// parse the a tag to collect the links
		if node.Type == html.ElementNode && node.Data == "a" {
			parseATag(node, links)
		}

		// if there is no child, then this is the end of the branch.
		if node.FirstChild == nil {
			// only collect actual text content that is trimmed of white space
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

// Parses a dom node to retrieve the href and text of an <a> tag
func parseATag(node *html.Node, links *[]Link) {
	link := Link{}
	link.Href = getHref(node.Attr)
	link.Text = parseNode(node.FirstChild, links)
	*links = append(*links, link)
}

// pulls the href attribute of of an <a> tag
func getHref(attrs []html.Attribute) string {
	var href string
	for _, attr := range attrs {
		if attr.Key == "href" {
			href = attr.Val
			break
		}
	}
	return href
}
