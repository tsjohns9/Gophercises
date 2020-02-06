package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

// JSONMap is a JSONMap
type JSONMap map[string]struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text,omitempty"`
		Arc  string `json:"arc,omitempty"`
	} `json:"options,omitempty"`
}

type handlerWithJSON struct {
	json JSONMap
}

func (h handlerWithJSON) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.String()[1:]
	fmt.Println(path)
	tmpl, err := template.ParseFiles("./index.html")
	if err != nil {
		res.Write([]byte("Failed to load the file!"))
		return
	}

	if keyValue, isInMap := h.json[path]; isInMap {
		tmpl.Execute(res, keyValue)
	} else {
		tmpl, err := template.ParseFiles("./error.html")
		if err != nil {
			res.Write([]byte("Failed to load the file!"))
			return
		}
		tmpl.Execute(res, "404 file not found")
	}
}

func main() {
	file := flag.String("file", "story.json", "JSON file for the adventure")
	flag.Parse()

	parsedJSON := openFile(*file)
	handler := handlerWithJSON{
		json: parsedJSON,
	}
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}

func openFile(fileName string) JSONMap {
	var pages JSONMap

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &pages)
	return pages
}
