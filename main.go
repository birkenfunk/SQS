package main

import (
	"flag"
	"log"
	"os"

	"codeberg.org/Birkenfunk/SQS/consts"
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
	println(GetHelloWorld())
	println(consts.GetWeatherServiceURL())
}

func GetHelloWorld() string {
	return "Hello, World!"
}

func IsTestRun() bool {
	return flag.Lookup("test.v").Value.(flag.Getter).Get().(bool)
}
