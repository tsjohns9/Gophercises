package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type yamlOutput []struct {
	Path string
	URL  string
}

type jsonOutput []struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.String()
		if val, isInMap := pathsToUrls[path]; isInMap {
			fmt.Println(val)
			http.Redirect(res, req, val, 301)
			return
		}
		fallback.ServeHTTP(res, req)
		return
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var parsedYaml yamlOutput
	err := yaml.Unmarshal(yml, &parsedYaml)
	if err != nil {
		return nil, err
	}
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.String()
		for _, urlMap := range parsedYaml {
			if urlMap.Path == path {
				http.Redirect(res, req, urlMap.URL, 301)
				return
			}
		}
		fallback.ServeHTTP(res, req)
		return
	}), nil
}

// JSONHandler is a json handler
func JSONHandler(js []byte, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.String()

		var output jsonOutput

		err := json.Unmarshal(js, &output)
		if err != nil {
			panic(err)
		}
		for _, obj := range output {
			if obj.Path == path {
				http.Redirect(res, req, obj.URL, 301)
				return
			}
		}
		fallback.ServeHTTP(res, req)
		return
	})
}
