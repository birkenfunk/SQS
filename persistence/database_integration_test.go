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
	db      IDatabase
	dto     *dtos.WeatherDto
	channel chan *dtos.WeatherDto
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
	suite.channel = make(chan *dtos.WeatherDto)
	weatherAddChannel = suite.channel
	pool = newPool()
}

func (suite *DatabaseIntegrationSuite) TearDownTest() {
	// cleanup
	_, err := pool.Get().Do("FLUSHALL")
	if err != nil {
		suite.FailNow(err.Error())
	}
	_ = pool.Close()
	close(suite.channel)
}

func (suite *DatabaseIntegrationSuite) TestNewDatabase() {
	db := NewDatabase()
	suite.NotNil(db)
}

func (suite *DatabaseIntegrationSuite) TestAddWeather() {
	// start the consumer
	go startWeatherConsumer()

	// add the weather to the channel
	suite.channel <- suite.dto

	// wait for the consumer to process the weather
	time.Sleep(1 * time.Second)

	result, err := suite.db.GetWeatherByLocation(suite.dto.Location)
	suite.Require().NoError(err)
	suite.Equal(suite.dto, result)
}

func (suite *DatabaseIntegrationSuite) TestAddWeather_Error_No_Time() {
	testWeather := &dtos.WeatherDto{
		Location:    "Berlin",
		Temperature: "20°C",
		Humidity:    "50%",
		SunHours:    8,
		WindSpeed:   "5m/s",
		Weather:     "Sunny",
	}
	go startWeatherConsumer()
	suite.channel <- testWeather
	// wait for the consumer to process the weather
	time.Sleep(1 * time.Second)
	result, err := suite.db.GetWeatherByLocation("Berlin")
	suite.Require().NoError(err)
	suite.Nil(result)
}

func (suite *DatabaseIntegrationSuite) TestAddWeather_Adding_Entry_Double() {
	go startWeatherConsumer()
	suite.channel <- suite.dto
	// wait for the consumer to process the weather
	time.Sleep(1 * time.Second)
	result, err := suite.db.GetWeatherByLocation(suite.dto.Location)
	suite.Require().NoError(err)
	suite.Equal(suite.dto, result)
	suite.channel <- suite.dto
	suite.Require().NoError(err)
	result, err = suite.db.GetWeatherByLocation(suite.dto.Location)
	suite.Require().NoError(err)
	suite.Equal(suite.dto, result)
}

func (suite *DatabaseIntegrationSuite) TestAddWeather_Double_Different_Values() {
	go startWeatherConsumer()
	suite.channel <- suite.dto
	// wait for the consumer to process the weather
	time.Sleep(1 * time.Second)
	suite.dto.Temperature = "25°C"
	suite.channel <- suite.dto
	// wait for the consumer to process the weather
	time.Sleep(1 * time.Second)
	result, err := suite.db.GetWeatherByLocation(suite.dto.Location)
	suite.Require().NoError(err)
	suite.Equal(suite.dto, result)
}

func (suite *DatabaseIntegrationSuite) TestGetWeatherByLocation() {
	processWeather(suite.dto)
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
	_, err := pool.Get().Do("SET", "Berlin", "not a json")
	suite.Require().NoError(err)
	result, err := suite.db.GetWeatherByLocation("Berlin")
	suite.Require().Error(err)
	suite.Nil(result)
}
