package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"./parse"
)

func main() {
	fileName := flag.String("file", "", "file to parse")
	url := flag.String("url", "", "url to html file to parse")
	flag.Parse()

	var r io.Reader

	// read from file
	if len(*fileName) > 0 {
		content := readFile(*fileName)
		r = strings.NewReader(content)
	}

	// fetch from url
	if len(*url) > 0 {
		data, err := fetch(*url)
		if err != nil {
			panic(err)
		}
		// ioutil.WriteFile("", data, 0666)
		r = strings.NewReader(string(data))
	}

	links, e := parse.Parse(r)

	if e != nil {
		panic(e)
	}

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
