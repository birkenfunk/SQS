package persistence

import (
	"os"
	"testing"
	"time"

	"codeberg.org/Birkenfunk/SQS/consts"
	"codeberg.org/Birkenfunk/SQS/dtos"
	"github.com/stretchr/testify/suite"
)

type DatabaseIntegrationSuite struct {
	suite.Suite
	db  IDatabase
	dto *dtos.WeatherDto
}

func TestDatabaseIntegrationSuite(t *testing.T) {
	suite.Run(t, new(DatabaseIntegrationSuite))
}

func (suite *DatabaseIntegrationSuite) SetupSuite() {
	redisURL, ok := os.LookupEnv("REDIS_ADDR")
	if !ok {
		redisURL = "localhost:6379"
	}
	consts.SetDBURL(redisURL)
}

func (suite *DatabaseIntegrationSuite) SetupTest() {
	suite.db = NewDatabase()
	suite.dto = &dtos.WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "50%",
		SunHours:    8,
		WindSpeed:   "5m/s",
		Weather:     "Sunny",
		Date:        time.Now().Format("2006-01-02"),
	}
}

func (suite *DatabaseIntegrationSuite) TearDownTest() {
	// cleanup
	_, err := (*suite.db.(*Database).con).Do("FLUSHALL")
	if err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *DatabaseIntegrationSuite) TestNewDatabase() {
	db := NewDatabase()
	suite.NotNil(db)
}

func (suite *DatabaseIntegrationSuite) TestAddWeather() {
	err := suite.db.AddWeather(suite.dto)
	suite.Require().NoError(err)
	result, err := suite.db.GetWeatherByLocation(suite.dto.Location)
	suite.Require().NoError(err)
	suite.Equal(suite.dto, result)
}

func (suite *DatabaseIntegrationSuite) TestAddWeather_Error_No_Time() {
	err := suite.db.AddWeather(&dtos.WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "50%",
		SunHours:    8,
		WindSpeed:   "5m/s",
		Weather:     "Sunny",
	})
	suite.Require().Error(err)
	result, err := suite.db.GetWeatherByLocation("Berlin")
	suite.Require().NoError(err)
	suite.Nil(result)
}

func (suite *DatabaseIntegrationSuite) TestAddWeather_Adding_Entry_Double() {
	err := suite.db.AddWeather(suite.dto)
	suite.Require().NoError(err)
	result, err := suite.db.GetWeatherByLocation(suite.dto.Location)
	suite.Require().NoError(err)
	suite.Equal(suite.dto, result)
	err = suite.db.AddWeather(suite.dto)
	suite.Require().NoError(err)
	result, err = suite.db.GetWeatherByLocation(suite.dto.Location)
	suite.Require().NoError(err)
	suite.Equal(suite.dto, result)
}

func (suite *DatabaseIntegrationSuite) TestAddWeather_Double_Different_Values() {
	err := suite.db.AddWeather(suite.dto)
	suite.Require().NoError(err)
	suite.dto.Temperature = "25°C"
	err = suite.db.AddWeather(suite.dto)
	suite.Require().NoError(err)
	result, err := suite.db.GetWeatherByLocation(suite.dto.Location)
	suite.Require().NoError(err)
	suite.Equal(suite.dto, result)
}

func (suite *DatabaseIntegrationSuite) TestGetWeatherByLocation() {
	err := suite.db.AddWeather(suite.dto)
	suite.Require().NoError(err)
	result, err := suite.db.GetWeatherByLocation("Berlin")
	suite.Require().NoError(err)
	suite.Equal(suite.dto, result)
}

func (suite *DatabaseIntegrationSuite) TestGetWeatherByLocation_NotFound() {
	result, err := suite.db.GetWeatherByLocation("Berlin")
	suite.Require().NoError(err)
	suite.Nil(result)
}

func (suite *DatabaseIntegrationSuite) TestGetWeatherByLocation_Error_in_Json() {
	db := suite.db.(*Database).con
	_, err := (*db).Do("SET", "Berlin", "not a json")
	suite.Require().NoError(err)
	result, err := suite.db.GetWeatherByLocation("Berlin")
	suite.Require().Error(err)
	suite.Nil(result)
}
