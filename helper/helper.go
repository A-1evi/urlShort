package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathTourls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if url, ok := pathTourls[req.URL.Path]; ok {
			http.Redirect(res, req, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(res, req)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func YamlHandler(yamlByte []byte, fallback http.Handler) (http.HandlerFunc, error) {
	//Parse the YAML data
	// and create a map of paths to urls
	// Callin MapHandler with the parsed YAML data
	pathUrls, err := parseYaml(yamlByte)
	if err != nil {
		return nil, err
	}
	pathToUrls := mapBuilder(pathUrls)
	return MapHandler(pathToUrls, fallback), nil
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func mapBuilder(pathUrls []pathUrl) map[string]string {
	pathToUrls := map[string]string{}
	for _, pathUrl := range pathUrls {
		pathToUrls[pathUrl.Path] = pathUrl.Url
	}

	return pathToUrls
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
