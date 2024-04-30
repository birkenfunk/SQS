package logic

import (
	"codeberg.org/Birkenfunk/SQS/dtos"
	"codeberg.org/Birkenfunk/SQS/persistence"
	"codeberg.org/Birkenfunk/SQS/service"
	"github.com/rs/zerolog/log"
)

type IWeather interface {
	GetWeather(location string) *dtos.WeatherDto
}

type Weather struct {
	weatherService service.IWeatherService
	database       persistence.IDatabase
}

func NewWeather() IWeather {
	return &Weather{
		weatherService: service.NewWeatherService(),
		database:       persistence.NewDatabase(),
	}
}

func (w *Weather) GetWeather(location string) *dtos.WeatherDto {
	/*result, err := w.database.GetWeatherByLocation(location)
	if err != nil {
		log.Err(err).Msg("Failed to get weather from database")
	}
	if result != nil {
		return result
	}*/
	result, err := w.weatherService.GetWeather(location)
	if err != nil {
		log.Err(err).Msg("Failed to get weather")
		return nil
	}
	persistence.GetWeatherAddChannel() <- result
	return result
}
