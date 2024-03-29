package main

import (
	"log"
	"os"

	"codeberg.org/Birkenfunk/SQS/consts"
	"codeberg.org/Birkenfunk/SQS/service"
	"github.com/joho/godotenv"
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
		log.Fatal(err)
	}
	consts.SetWeatherServiceURL(os.Getenv("WEATHER_SERVICE_API_URL"))
}

func main() {
	var ws service.IWeatherService = service.WeatherService{}
	// Check if the weather service is available
	err := ws.GetHealth()
	if err != nil {
		log.Fatal(err)
	}
	weather, err := ws.GetWeather("Berlin")
	if err != nil {
		log.Fatal(err)
	}
	println(weather.String())
}

func GetHelloWorld() string {
	return "Hello, World!"
}
