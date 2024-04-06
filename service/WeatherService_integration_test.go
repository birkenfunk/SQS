package service

import (
	"os"
	"testing"

	"codeberg.org/Birkenfunk/SQS/consts"
	"codeberg.org/Birkenfunk/SQS/dtos"
	"github.com/stretchr/testify/suite"
)

type WeatherServiceSuite struct {
	suite.Suite
	externalAPI string
	ws          *WeatherService
}

func TestWeatherServiceSuite(t *testing.T) {
	suite.Run(t, new(WeatherServiceSuite))
}

func (wss *WeatherServiceSuite) SetupSuite() {
	externalAPI, ok := os.LookupEnv("EXTERNAL_API")
	if !ok {
		externalAPI = "http://localhost:3000"
	}
	wss.externalAPI = externalAPI
	wss.ws = &WeatherService{}
}

// TestGetHealth_Success tests the GetHealth function with a successful external API
// If the Test Fails you should check if the external API is running under localhost:3000 or the URL you specified in the EXTERNAL_API environment variable
func (wss *WeatherServiceSuite) TestGetHealth_Success() {
	// given:
	consts.SetWeatherServiceURL(wss.externalAPI)

	// when:
	err := wss.ws.GetHealth()

	// then:
	wss.NoError(err)
}

// TestGetHealth_Fail tests the GetHealth function with a failing external API
// If the Test Fails you should check if the external API is running under localhost:3000 or the URL you specified in the EXTERNAL_API environment variable
func (wss *WeatherServiceSuite) TestGetHealth_Fail() {
	// given:
	consts.SetWeatherServiceURL(wss.externalAPI + "1")

	// when:
	err := wss.ws.GetHealth()

	// then:
	wss.Error(err)
}

// TestGetWeather_Success tests the GetWeather function with a successful external API
// If the Test Fails you should check if the external API is running under localhost:3000 or the URL you specified in the EXTERNAL_API environment variable
func (wss *WeatherServiceSuite) TestGetWeather_Success() {
	// given:
	consts.SetWeatherServiceURL(wss.externalAPI)

	// when:
	weather, err := wss.ws.GetWeather("test")

	// then:
	wss.Require().NoError(err)
	wss.Equal(&dtos.WeatherDto{
		Location:    "Test",
		Temperature: "20°C",
		Humidity:    "50%",
		SunHours:    5,
		WindSpeed:   "10km/h",
		Weather:     "Sunny",
		Date:        "2024-01-01",
	}, weather)
}

// TestGetWeather_Fail tests the GetWeather function with a failing external API
// If the Test Fails you should check if the external API is running under localhost:3000 or the URL you specified in the EXTERNAL_API environment variable
func (wss *WeatherServiceSuite) TestGetWeather_Fail() {
	// given:
	consts.SetWeatherServiceURL(wss.externalAPI + "1")

	// when:
	weather, err := wss.ws.GetWeather("test")

	// then:
	wss.Require().Error(err)
	wss.Nil(weather)
}
