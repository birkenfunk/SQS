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
	weather      IWeather
	weatherMock  *mocks.IWeatherService
	databaseMock *mocks.IDatabase
	weatherDto   dtos.WeatherDto
}

func TestWeatherSuite(t *testing.T) {
	suite.Run(t, &WeatherSuite{})
}

func (suite *WeatherSuite) SetupTest() {
	suite.weatherMock = new(mocks.IWeatherService)
	suite.databaseMock = new(mocks.IDatabase)
	suite.weather = &Weather{suite.weatherMock, suite.databaseMock}
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

func (suite *WeatherSuite) TestGetWeather_Success_Not_In_DB() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(&suite.weatherDto, nil)
	suite.databaseMock.On("GetWeatherByLocation", "Berlin").Return(nil, nil)
	suite.databaseMock.On("AddWeather", &suite.weatherDto).Return(nil)
	result := suite.weather.GetWeather("Berlin")
	suite.Equal(&suite.weatherDto, result)
	suite.databaseMock.AssertCalled(suite.T(), "GetWeatherByLocation", "Berlin")
	suite.databaseMock.AssertCalled(suite.T(), "AddWeather", &suite.weatherDto)
	suite.weatherMock.AssertCalled(suite.T(), "GetWeather", "Berlin")
}

func (suite *WeatherSuite) TestGetWeather_Success_err_from_database() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(&suite.weatherDto, nil)
	suite.databaseMock.On("GetWeatherByLocation", "Berlin").Return(nil, err)
	suite.databaseMock.On("AddWeather", &suite.weatherDto).Return(nil)
	result := suite.weather.GetWeather("Berlin")
	suite.Equal(&suite.weatherDto, result)
	suite.databaseMock.AssertCalled(suite.T(), "GetWeatherByLocation", "Berlin")
	suite.weatherMock.AssertCalled(suite.T(), "GetWeather", "Berlin")
	suite.databaseMock.AssertCalled(suite.T(), "AddWeather", &suite.weatherDto)
}

func (suite *WeatherSuite) TestGetWeather_Success_In_DB() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(nil, nil)
	suite.databaseMock.On("GetWeatherByLocation", "Berlin").Return(&suite.weatherDto, nil)
	result := suite.weather.GetWeather("Berlin")
	suite.Equal(&suite.weatherDto, result)
	suite.databaseMock.AssertCalled(suite.T(), "GetWeatherByLocation", "Berlin")
	suite.weatherMock.AssertNotCalled(suite.T(), "GetWeather", "Berlin")
	suite.databaseMock.AssertNotCalled(suite.T(), "AddWeather", &suite.weatherDto)
}

func (suite *WeatherSuite) TestGetWeather_Fail_err_from_weatherService() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(nil, err)
	suite.databaseMock.On("GetWeatherByLocation", "Berlin").Return(nil, nil)
	suite.databaseMock.On("AddWeather", &suite.weatherDto).Return(nil)
	result := suite.weather.GetWeather("Berlin")
	suite.Nil(result)
	suite.databaseMock.AssertCalled(suite.T(), "GetWeatherByLocation", "Berlin")
	suite.weatherMock.AssertCalled(suite.T(), "GetWeather", "Berlin")
	suite.databaseMock.AssertNotCalled(suite.T(), "AddWeather", &suite.weatherDto)
}
