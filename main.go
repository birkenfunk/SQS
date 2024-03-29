package main

import (
	"codeberg.org/Birkenfunk/SQS/consts"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	env := os.Getenv("ENV")
	var err error
	if env == "test" {
		godotenv.Load("test.env")
	} else {
		godotenv.Load()
	}
	if err != nil {
		log.Fatal(err)
	}
	consts.SetWeatherServiceURL(os.Getenv("WEATHER_SERVICE_API_URL"))
	//write to file
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
