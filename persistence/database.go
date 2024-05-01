package persistence

import (
	"encoding/json"
	"time"

	"codeberg.org/Birkenfunk/SQS/consts"
	"codeberg.org/Birkenfunk/SQS/dtos"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
)

type IDatabase interface {
	GetWeatherByLocation(location string) (*dtos.WeatherDto, error)
}

var weatherAddChannel chan *dtos.WeatherDto

func GetWeatherAddChannel() chan *dtos.WeatherDto {
	return weatherAddChannel
}

var pool redis.Pool

func startWeatherConsumer() {
	for weather := range GetWeatherAddChannel() {
		processWeather(weather)
	}
}

func InitDB() {
	weatherAddChannel = make(chan *dtos.WeatherDto, 1000000)
	pool = newPool()
	go startWeatherConsumer()
}

func newPool() redis.Pool {
	return redis.Pool{
		Dial:      func() (redis.Conn, error) { return redis.Dial("tcp", consts.GetDBURL()) },
		MaxIdle:   500,
		MaxActive: 50000,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func processWeather(weather *dtos.WeatherDto) {
	log.Debug().Msg("Adding weather for " + weather.Location + "to redis")
	expTime, err := time.Parse("2006-01-02 15:04:05", weather.Date+" 23:59:59")
	if err != nil {
		log.Error().Err(err).Msg("Could not parse date")
	}
	weatherJSON, err := json.Marshal(weather)
	if err != nil {
		log.Error().Err(err).Msg("Could not marshal weather")
	}
	con := pool.Get()
	_, err = con.Do("SET", weather.Location, weatherJSON, "EXAT", expTime.Unix())
	err = con.Close()
	if err != nil {
		log.Error().Err(err).Msg("Could not close redis connection")
	}
	if err != nil {
		log.Error().Err(err).Msg("Could not add weather to redis")
	}
}

type Database struct {
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) GetWeatherByLocation(location string) (*dtos.WeatherDto, error) {
	con := pool.Get()
	result, err := con.Do("GET", location)
	if err != nil {
		log.Error().Err(err).Msg("Could not get weather from redis")
		return nil, err
	}
	err = con.Close()
	if err != nil {
		log.Error().Err(err).Msg("Could not close redis connection")
	}
	if result == nil {
		return nil, nil
	}
	var weather dtos.WeatherDto
	err = json.Unmarshal(result.([]byte), &weather)
	if err != nil {
		log.Error().Err(err).Msg("Could not unmarshal weather")
		return nil, err
	}
	return &weather, nil
}

// SetWeatherAddChannel sets the channel for adding weather to the database
func SetWeatherAddChannel(channel chan *dtos.WeatherDto) {
	weatherAddChannel = channel
}

// StartWeatherConsumer starts the consumer for adding weather to the database
// For testing purposes only
func StartWeatherConsumer() {
	go startWeatherConsumer()
}
