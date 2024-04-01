package logic

import (
	"codeberg.org/Birkenfunk/SQS/dtos"
	"codeberg.org/Birkenfunk/SQS/service"
	"github.com/rs/zerolog/log"
)

var weatherService service.IWeatherService = &service.WeatherService{}

type IWeather interface {
	GetWeather(location string) *dtos.WeatherDto
}

type Weather struct{}

func (w *Weather) GetWeather(location string) *dtos.WeatherDto {
	result, err := weatherService.GetWeather(location)
	if err != nil {
		log.Err(err).Msg("Failed to get weather")
		return nil
	}
	return result
}
