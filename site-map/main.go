package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"../link-parser/parse"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

// Link is a link from the parser package
type Link = parse.Link

func main() {
	url := flag.String("url", "", "url to html file to parse")
	depth := flag.Int("depth", 2, "max depth")
	flag.Parse()

	if *url == "" {
		fmt.Println("A url is required")
		return
	}

	// returns an array of all urls collected from the initial url
	// up to the specified depth
	urls := bfs(removeTrailingSlash(*url), *depth)

	var locations []loc

	// convert the url string into a struct for parsing into xml
	for _, url := range urls {
		locations = append(locations, loc{Value: url})
	}

	// generate the xml
	x, e := xml.Marshal(urlset{Urls: locations, Xmlns: xmlns})

	if e != nil {
		panic(e)
	}
	ioutil.WriteFile("sitemap.xml", x, 0666)
}

func bfs(domain string, depth int) []string {
	seen := make(map[string]string)
	notSeen := make(map[string]string)
	notSeen[domain] = domain
	for i := 0; i < depth; i++ {
		for _, url := range notSeen {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = url
			for _, link := range fetchAndParse(url, domain) {
				notSeen[link] = removeTrailingSlash(link)
			}
		}
	}
	var arr []string
	for url := range seen {
		arr = append(arr, url)
	}
	return arr
}

func filter(links []Link, keepItem func(l Link) bool) []string {
	var urls []string
	for _, link := range links {
		if keepItem(link) {
			urls = append(urls, link.Href)
		}

	}
	return urls
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, err
	}
	return bytes, nil
}

func isValidLink(prefix string) func(Link) bool {
	return func(l Link) bool {
		return strings.HasPrefix(l.Href, prefix)
	}
}

func fetchAndParse(domain string, prefix string) []string {
	bytes, err := fetch(domain)
	if err != nil {
		panic(err)
	}

	// read in the html content
	r := strings.NewReader(string(bytes))

	// get all <a> tags from the reader
	links, err := parse.Parse(r)
	if err != nil {
		panic(err)
	}

	filteredLinks := filter(setPrefix(links, prefix), isValidLink(prefix))
	return filteredLinks
}

func setPrefix(links []Link, prefix string) []Link {
	var withPrefix []Link
	for _, link := range links {
		if strings.HasPrefix(link.Href, "//") {
			continue
		}
		if strings.HasPrefix(link.Href, "/") {
			link.Href = prefix + link.Href
			withPrefix = append(withPrefix, link)
		}
	}
	return withPrefix
}

func removeTrailingSlash(domain string) string {
	if strings.HasSuffix(domain, "/") {
		domain = domain[:len(domain)-1]
	}
	return domain
}
