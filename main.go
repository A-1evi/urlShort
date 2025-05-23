package main

import (
	"fmt"
	"net/http"
	urlshort "urlShortner/helper"
)

func main() {
	mux := defaultMux()
	pathToUrls := map[string]string{
		"/urlshortdocs": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godocs":  "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, err := urlshort.YamlHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux

}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Helllo world!")
}
