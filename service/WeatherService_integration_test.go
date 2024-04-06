package service

import (
	"os"
	"testing"

	"codeberg.org/Birkenfunk/SQS/consts"
	"codeberg.org/Birkenfunk/SQS/dtos"
	"github.com/stretchr/testify/suite"
)

type WeatherServiceIntegationSuite struct {
	suite.Suite
	externalAPI string
	ws          *WeatherService
}

func TestWeatherServiceSuite(t *testing.T) {
	suite.Run(t, new(WeatherServiceIntegationSuite))
}

func (wsis *WeatherServiceIntegationSuite) SetupSuite() {
	externalAPI, ok := os.LookupEnv("EXTERNAL_API")
	if !ok {
		externalAPI = "http://localhost:3000"
	}
	wsis.externalAPI = externalAPI
	wsis.ws = &WeatherService{}
}

// TestGetHealth_Success tests the GetHealth function with a successful external API
// If the Test Fails you should check if the external API is running under localhost:3000 or the URL you specified in the EXTERNAL_API environment variable
func (wsis *WeatherServiceIntegationSuite) TestGetHealth_Success() {
	// given:
	consts.SetWeatherServiceURL(wsis.externalAPI)

	// when:
	err := wsis.ws.GetHealth()

	// then:
	wsis.NoError(err)
}

// TestGetHealth_Fail tests the GetHealth function with a failing external API
// If the Test Fails you should check if the external API is running under localhost:3000 or the URL you specified in the EXTERNAL_API environment variable
func (wsis *WeatherServiceIntegationSuite) TestGetHealth_Fail() {
	// given:
	consts.SetWeatherServiceURL(wsis.externalAPI + "1")

	// when:
	err := wsis.ws.GetHealth()

	// then:
	wsis.Error(err)
}

// TestGetWeather_Success tests the GetWeather function with a successful external API
// If the Test Fails you should check if the external API is running under localhost:3000 or the URL you specified in the EXTERNAL_API environment variable
func (wsis *WeatherServiceIntegationSuite) TestGetWeather_Success() {
	// given:
	consts.SetWeatherServiceURL(wsis.externalAPI)

	// when:
	weather, err := wsis.ws.GetWeather("test")

	// then:
	wsis.Require().NoError(err)
	wsis.Equal(&dtos.WeatherDto{
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
func (wsis *WeatherServiceIntegationSuite) TestGetWeather_Fail() {
	// given:
	consts.SetWeatherServiceURL(wsis.externalAPI + "1")

	// when:
	weather, err := wsis.ws.GetWeather("test")

	// then:
	wsis.Require().Error(err)
	wsis.Nil(weather)
}
