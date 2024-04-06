package logic

import (
	"fmt"
	"testing"

	"codeberg.org/Birkenfunk/SQS/dtos"
	"codeberg.org/Birkenfunk/SQS/mocks"
	"github.com/stretchr/testify/suite"
)

type WeatherSuite struct {
	suite.Suite
	weather     IWeather
	weatherMock *mocks.IWeatherService
	weatherDto  dtos.WeatherDto
}

func TestWeatherSuite(t *testing.T) {
	suite.Run(t, &WeatherSuite{})
}

func (suite *WeatherSuite) SetupTest() {
	suite.weatherMock = new(mocks.IWeatherService)
	suite.weather = NewWeather()
	suite.weather.(*Weather).weatherService = suite.weatherMock
}

func (suite *WeatherSuite) SetupSuite() {
	suite.weatherDto = dtos.WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "20%",
		SunHours:    5,
		WindSpeed:   "50m/s",
		Weather:     "Sunny",
		Date:        "2021-09-01",
	}
}

var err = fmt.Errorf("failed to get weather")

func (suite *WeatherSuite) TestGetWeather_Success() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(&suite.weatherDto, nil)
	result := suite.weather.GetWeather("Berlin")
	suite.Equal(&suite.weatherDto, result)
}

func (suite *WeatherSuite) TestGetWeather_Fail() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(nil, err)
	result := suite.weather.GetWeather("Berlin")
	suite.Nil(result)
}
