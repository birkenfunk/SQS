package testfixtures

import (
	"os"

	"codeberg.org/Birkenfunk/SQS/consts"
)

func SetUpAllVariables() {
	SetupWeatherServiceURL()
	SetupPort()
	SetupDBURL()
}

func SetupWeatherServiceURL() {
	weatherServiceURL, ok := os.LookupEnv("EXTERNAL_API")
	if !ok {
		weatherServiceURL = "http://localhost:3000"
	}
	consts.SetWeatherServiceURL(weatherServiceURL)
}

func SetupPort() {
	consts.SetPortFromString("3000")
}

func SetupDBURL() {
	dbURL, ok := os.LookupEnv("REDIS_ADDR")
	if !ok {
		dbURL = "localhost:6379"
	}
	consts.SetDBURL(dbURL)
}
