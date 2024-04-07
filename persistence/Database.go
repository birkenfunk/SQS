package persistence

import (
	"encoding/json"
	"time"

	"codeberg.org/Birkenfunk/SQS/consts"
	"codeberg.org/Birkenfunk/SQS/dtos"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
)

type Database struct {
	con *redis.Conn
}

func NewDatabase() *Database {
	newCon, err := redis.Dial("tcp", consts.GetDBURL())
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to redis")
	}
	_, err = newCon.Do("PING")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to redis")
		return nil
	}
	return &Database{
		con: &newCon,
	}
}

func (db *Database) AddWeather(dto *dtos.WeatherDto) error {
	log.Debug().Msg("Adding weather for " + dto.Location + "to redis")
	expTime, err := time.Parse("2006-01-02 15:04:05", dto.Date+" 23:59:59")
	if err != nil {
		log.Error().Err(err).Msg("Could not parse date")
		return err
	}
	weatherJson, err := json.Marshal(dto)
	if err != nil {
		log.Error().Err(err).Msg("Could not marshal weather")
		return err
	}
	_, err = (*db.con).Do("SET", dto.Location, weatherJson, "EXAT", expTime.Unix())
	if err != nil {
		log.Error().Err(err).Msg("Could not add weather to redis")
		return err
	}
	return nil
}

func (db *Database) GetWeatherByLocation(location string) (*dtos.WeatherDto, error) {
	result, err := (*db.con).Do("GET", location)
	if err != nil {
		log.Error().Err(err).Msg("Could not get weather from redis")
		return nil, err
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
