package main

import (
	"net/http"
	"os"

	"codeberg.org/Birkenfunk/SQS/consts"
	"codeberg.org/Birkenfunk/SQS/presentation"
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
		log.Fatal().Err(err)
	}
	consts.SetWeatherServiceURL(os.Getenv("WEATHER_SERVICE_API_URL"))
	consts.SetPortFromString(os.Getenv("PORT"))
}

func main() {
	router := presentation.NewRouter()
	routes := router.InitRouter()

	// Start the server
	err := http.ListenAndServe(":"+consts.GetPort(), routes)
	if err != nil {
		log.Fatal().Err(err)
	}
}

func GetHelloWorld() string {
	return "Hello, World!"
}
