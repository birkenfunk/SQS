package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"strings"

	"codeberg.org/Birkenfunk/SQS/presentation"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/go-chi/docgen/raml"
	"github.com/rs/zerolog/log"
)

func main() {
	router := presentation.NewRouter()
	routes := router.InitRouter()

	// Save the routes to a file
	if err := os.Remove("routes.raml"); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatal().Err(err)
	}

	f, err := os.Create("routes.raml")

	if err != nil {
		log.Fatal().Err(err)
	}

	defer f.Close()

	ramlDocs := &raml.RAML{
		Title:     "RAML Representation of RESTful API",
		BaseUri:   "http://localhost:8080/api/v1",
		Version:   "v1.0",
		MediaType: "application/json",
	}

	if err := chi.Walk(routes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		handlerInfo := docgen.GetFuncInfo(handler)
		resource := &raml.Resource{
			DisplayName: strings.ToUpper(method) + " " + route,
			Description: "Handler Function: " + handlerInfo.Func + "\nComment: " + handlerInfo.Comment,
		}

		return ramlDocs.Add(method, route, resource)
	}); err != nil {
		log.Fatal().Msgf("error: %v", err)
	}

	raml, err := yaml.Marshal(ramlDocs)

	if err != nil {
		log.Fatal().Err(err)
	}

	if _, err = f.Write(append([]byte("#%RAML 1.0\n---\n"), raml...)); err != nil {
		log.Fatal().Err(err)
	}
}
