package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/go-chi/docgen/raml"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MainSuite struct {
	suite.Suite
}

func TestMainSuite(t *testing.T) {
	suite.Run(t, &MainSuite{})
}

func (ms *MainSuite) TearDownTest() {
	http.DefaultClient.CloseIdleConnections()
}

func (ms *MainSuite) TestStartServer() {
	// given:
	handler := http.Handler(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	// when:

	wg := &sync.WaitGroup{}
	server := startServer(handler, wg)

	// then:
	ms.Require().NotNil(server)
	ms.Require().NoError(server.Shutdown(context.TODO()))
}

func (ms *MainSuite) TestStartServer_Error() {
	// given:
	handler := http.Handler(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	// when:

	wg := &sync.WaitGroup{}
	server := startServer(handler, wg)
	wg2 := &sync.WaitGroup{}
	server2 := startServer(handler, wg2)

	// then:
	ms.Require().NotNil(server)
	ms.Require().NotNil(server2)
	ms.Require().NoError(server.Shutdown(context.TODO()))
	ms.Require().NoError(server2.Shutdown(context.TODO()))
}

func (ms *MainSuite) TestInitRoutes() {
	// when:
	routes := initRoutes()

	// then:
	ms.Require().NotNil(routes)
}

func (ms *MainSuite) TestGenerateRAML() {
	routes := initRoutes()

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
