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
	parseNode(document)

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

func parseNode(node *html.Node) string {
	if node != nil {
		if node.FirstChild == nil {
			x := strings.TrimSpace(node.Data)
			if len(x) > 0 {
				return x
			}
		}
		if node.NextSibling != nil {
			parseNode(node.NextSibling)
		}
		if node.FirstChild != nil {
			if node.Data == "a" {
				parseLink(node)
			} else {
				parseNode(node.FirstChild)
			}
		}
	}

	return ""
}

func parseLink(node *html.Node) {
	Href := getHref(node.Attr)
	Text := parseNode(node.FirstChild)
	link := Link{
		Href,
		Text,
	}
	// fmt.Println("Link:", link)
	Links = append(Links, link)
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

// Ex: FirstChild == html
// FirstChild.Data == head
// FirstChild.NextSibling.Data == body
