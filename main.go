package main

import (
	"log"
	"os"

	"codeberg.org/Birkenfunk/SQS/consts"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	consts.SetWeatherServiceURL(os.Getenv("WEATHER_SERVICE_API_URL"))
}

func main() {
	println(GetHelloWorld())
	println(consts.GetWeatherServiceURL())
}

func GetHelloWorld() string {
	return "Hello, World!"
}
