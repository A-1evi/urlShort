package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	urlshort "urlShortner/helper"
)

func main() {
	yamlFileName := flag.String("yml", "path.yml", "this is yaml file containing path of of url")

	flag.Parse()
	yamlFile, err := os.ReadFile(*yamlFileName)
	if err != nil {
		fmt.Println("Error reading yaml file:", err)
		return
	}

	mux := defaultMux()
	pathToUrls := map[string]string{
		"/urlshortdocs": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godocs":  "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathToUrls, mux)

	// Use the yamlFile content instead of hardcoded yaml string
	yamlHandler, err := urlshort.YamlHandler(yamlFile, mapHandler)
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
	fmt.Fprintf(res, "Hello world!")
}
