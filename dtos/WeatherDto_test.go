package dtos

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WeatherDtoSuite struct {
	suite.Suite
	dto     *WeatherDto
	dtoSame *WeatherDto
	dtoDiff *WeatherDto
}

func TestWeatherDtoSuite(t *testing.T) {
	suite.Run(t, &WeatherDtoSuite{})
}

func (suite *WeatherDtoSuite) SetupTest() {
	suite.dto = &WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "50%",
		SunHours:    8,
		WindSpeed:   "5m/s",
		Weather:     "Sunny",
		Date:        "2021-09-01",
	}
	suite.dtoSame = &WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "50%",
		SunHours:    8,
		WindSpeed:   "5m/s",
		Weather:     "Sunny",
		Date:        "2021-09-01",
	}
	suite.dtoDiff = &WeatherDto{
		Location:    "Paris",
		Temperature: "25°C",
		Humidity:    "60%",
		SunHours:    10,
		WindSpeed:   "3m/s",
		Weather:     "Cloudy",
		Date:        "2021-09-01",
	}
}

func (suite *WeatherDtoSuite) TestWeatherDto_Equals_True() {
	suite.True(suite.dto.Equals(suite.dtoSame))
}

func (suite *WeatherDtoSuite) TestWeatherDto_Equals_False() {
	suite.False(suite.dto.Equals(suite.dtoDiff))
}

func (suite *WeatherDtoSuite) TestWeatherDto_Equals_False_Nil() {
	suite.False(suite.dto.Equals(nil))
}

func (suite *WeatherDtoSuite) TestWeatherDto_String() {
	suite.Equal("Location: Berlin\nTemperature: 20°C\nHumidity: 50%\nSunHours: 8\nWindSpeed: 5m/s\nWeather: Sunny\nDate: 2021-09-01\n", suite.dto.String())
}

func (suite *WeatherDtoSuite) TestWeatherDto_Diff() {
	suite.Equal("Location: Berlin != Paris\nTemperature: 20°C != 25°C\nHumidity: 50% != 60%\nSunHours: 8 != 10\nWindSpeed: 5m/s != 3m/s\nWeather: Sunny != Cloudy\n", suite.dto.Diff(suite.dtoDiff))
}

func (suite *WeatherDtoSuite) TestWeatherDto_Diff_Same() {
	suite.Equal("", suite.dto.Diff(suite.dtoSame))
}

func (suite *WeatherDtoSuite) TestWeatherDto_Diff_Nil() {
	suite.Equal("nil", suite.dto.Diff(nil))
}
