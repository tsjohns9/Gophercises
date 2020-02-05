package main

import (
	"fmt"
	"net/http"

	"./urlshort"
)

// Build the MapHandler using the mux as the fallback
var pathsToUrls = map[string]string{
	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}

var jsonUrls = `[
	{
		"path": "/tsj",
		"url": "https://github.com/tsjohns9"
	},
	{
		"path": "/react",
		"url": "https://github.com/facebook/react"
	}
]`

// Build the YAMLHandler using the mapHandler as the fallback
var yaml = `
    - path: /urlshort
      url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
      url: https://github.com/gophercises/urlshort/tree/solution`

func main() {
	mux := defaultMux()

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	jsonHandler := urlshort.JSONHandler([]byte(jsonUrls), mapHandler)

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), jsonHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
