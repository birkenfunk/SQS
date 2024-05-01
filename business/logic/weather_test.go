package logic

import (
	"fmt"
	"testing"

	"codeberg.org/Birkenfunk/SQS/dtos"
	persistence2 "codeberg.org/Birkenfunk/SQS/mocks/codeberg.org/Birkenfunk/SQS/persistence"
	service2 "codeberg.org/Birkenfunk/SQS/mocks/codeberg.org/Birkenfunk/SQS/service"
	"codeberg.org/Birkenfunk/SQS/persistence"
	"codeberg.org/Birkenfunk/SQS/testfixtures"
	"github.com/stretchr/testify/suite"
)

type WeatherSuite struct {
	suite.Suite
	weather      IWeather
	weatherMock  *service2.MockIWeatherService
	databaseMock *persistence2.MockIDatabase
	weatherDto   dtos.WeatherDto
	chanel       chan *dtos.WeatherDto
}

func TestWeatherSuite(t *testing.T) {
	suite.Run(t, &WeatherSuite{})
}

func (suite *WeatherSuite) SetupTest() {
	suite.weatherMock = new(service2.MockIWeatherService)
	suite.databaseMock = new(persistence2.MockIDatabase)
	suite.weather = &Weather{suite.weatherMock, suite.databaseMock}
	suite.chanel = make(chan *dtos.WeatherDto)
	persistence.SetWeatherAddChannel(suite.chanel)
	persistence.StartWeatherConsumer()
}

func (suite *WeatherSuite) TearDownTest() {
	close(suite.chanel)
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

func (suite *WeatherSuite) TestNewWeather() {
	testfixtures.SetUpAllVariables()
	result := NewWeather()
	suite.NotNil(result)
	suite.IsType(&Weather{}, result)
}

func (suite *WeatherSuite) TestGetWeather_Success_Not_In_DB() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(&suite.weatherDto, nil)
	suite.databaseMock.On("GetWeatherByLocation", "Berlin").Return(nil, nil)
	result := suite.weather.GetWeather("Berlin")
	suite.Equal(&suite.weatherDto, result)
	suite.databaseMock.AssertCalled(suite.T(), "GetWeatherByLocation", "Berlin")
	suite.weatherMock.AssertCalled(suite.T(), "GetWeather", "Berlin")
}

func (suite *WeatherSuite) TestGetWeather_Success_err_from_database() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(&suite.weatherDto, nil)
	suite.databaseMock.On("GetWeatherByLocation", "Berlin").Return(nil, err)
	result := suite.weather.GetWeather("Berlin")
	suite.Equal(&suite.weatherDto, result)
	suite.databaseMock.AssertCalled(suite.T(), "GetWeatherByLocation", "Berlin")
	suite.weatherMock.AssertCalled(suite.T(), "GetWeather", "Berlin")
}

func (suite *WeatherSuite) TestGetWeather_Success_In_DB() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(nil, nil)
	suite.databaseMock.On("GetWeatherByLocation", "Berlin").Return(&suite.weatherDto, nil)
	result := suite.weather.GetWeather("Berlin")
	suite.Equal(&suite.weatherDto, result)
	suite.databaseMock.AssertCalled(suite.T(), "GetWeatherByLocation", "Berlin")
	suite.weatherMock.AssertNotCalled(suite.T(), "GetWeather", "Berlin")
}

func (suite *WeatherSuite) TestGetWeather_Success_err_adding_to_db() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(&suite.weatherDto, nil)
	suite.databaseMock.On("GetWeatherByLocation", "Berlin").Return(nil, nil)
	result := suite.weather.GetWeather("Berlin")
	suite.Equal(&suite.weatherDto, result)
	suite.databaseMock.AssertCalled(suite.T(), "GetWeatherByLocation", "Berlin")
	suite.weatherMock.AssertCalled(suite.T(), "GetWeather", "Berlin")
}

func (suite *WeatherSuite) TestGetWeather_Fail_err_from_weatherService() {
	suite.weatherMock.On("GetWeather", "Berlin").Return(nil, err)
	suite.databaseMock.On("GetWeatherByLocation", "Berlin").Return(nil, nil)
	result := suite.weather.GetWeather("Berlin")
	suite.Nil(result)
	suite.databaseMock.AssertCalled(suite.T(), "GetWeatherByLocation", "Berlin")
	suite.weatherMock.AssertCalled(suite.T(), "GetWeather", "Berlin")
}
