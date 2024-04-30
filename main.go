package main

import (
	"codeberg.org/Birkenfunk/SQS/persistence"
	"errors"
	"net/http"
	"os"
	"sync"

	"codeberg.org/Birkenfunk/SQS/consts"
	"codeberg.org/Birkenfunk/SQS/presentation"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	env := os.Getenv("ENV")
	var err error
	if env == "test" {
		err = godotenv.Load("test.env")
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		log.Error().Err(err)
	}
	consts.SetWeatherServiceURL(os.Getenv("WEATHER_SERVICE_API_URL"))
	consts.SetPortFromString(os.Getenv("PORT"))
	consts.SetDBURL(os.Getenv("REDIS_URL"))
	log.Debug().Msg("Initialized")
}

func main() {
	routes := initRoutes()

	persistence.InitDB()

	// Start the server
	wg := &sync.WaitGroup{}
	startServer(routes, wg)
	wg.Wait()
}

func initRoutes() *chi.Mux {
	router := presentation.NewRouter()
	routes := router.InitRouter()
	return routes
}

func startServer(handler http.Handler, wg *sync.WaitGroup) *http.Server {
	srv := &http.Server{
		Addr:    ":" + consts.GetPort(),
		Handler: handler,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info().Msgf("Starting server on port %s", consts.GetPort())
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Server failed")
		}
	}()
	return srv
}
