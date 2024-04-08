package main

import (
	"context"
	"net/http"
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
