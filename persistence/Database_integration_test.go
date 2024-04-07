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
	db IDatabase
}

func TestDatabaseIntegrationSuite(t *testing.T) {
	suite.Run(t, new(DatabaseIntegrationSuite))
}

func (suite *DatabaseIntegrationSuite) SetupSuite() {
	redisUrl, ok := os.LookupEnv("REDIS_ADDR")
	if !ok {
		redisUrl = "localhost:6379"
	}
	consts.SetDBURL(redisUrl)
}

func (suite *DatabaseIntegrationSuite) SetupTest() {
	suite.db = NewDatabase()
}

func (suite *DatabaseIntegrationSuite) TearDownTest() {
	// cleanup
	_, err := (*suite.db.(*Database).con).Do("FLUSHALL")
	if err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *DatabaseIntegrationSuite) TestAddWeather() {
	err := suite.db.AddWeather(&dtos.WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "50%",
		SunHours:    8,
		WindSpeed:   "5m/s",
		Weather:     "Sunny",
		Date:        "2021-09-01",
	})
	suite.NoError(err)
}

func (suite *DatabaseIntegrationSuite) TestAddWeather_Error_in_DTO() {
	err := suite.db.AddWeather(&dtos.WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "50%",
		SunHours:    8,
		WindSpeed:   "5m/s",
		Weather:     "Sunny",
	})
	suite.Error(err)
}

func (suite *DatabaseIntegrationSuite) TestGetWeatherByLocation() {
	dto := &dtos.WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "50%",
		SunHours:    8,
		WindSpeed:   "5m/s",
		Weather:     "Sunny",
		Date:        time.Now().Format("2006-01-02"),
	}
	err := suite.db.AddWeather(dto)
	suite.NoError(err)
	result, err := suite.db.GetWeatherByLocation("Berlin")
	suite.NoError(err)
	suite.Equal(dto, result)
}

func (suite *DatabaseIntegrationSuite) TestGetWeatherByLocation_NotFound() {
	result, err := suite.db.GetWeatherByLocation("Berlin")
	suite.NoError(err)
	suite.Nil(result)
}
