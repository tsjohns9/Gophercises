package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
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

func main() {
	file := flag.String("file", "story.json", "JSON file for the adventure")
	flag.Parse()

	openFile(*file)
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
