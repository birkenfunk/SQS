package logic

import (
	"codeberg.org/Birkenfunk/SQS/dtos"
	"codeberg.org/Birkenfunk/SQS/service"
	"github.com/rs/zerolog/log"
)

type IWeather interface {
	GetWeather(location string) *dtos.WeatherDto
}

type Weather struct {
	weatherService service.IWeatherService
}

func NewWeather() IWeather {
	return &Weather{weatherService: service.NewWeatherService()}
}

func (w *Weather) GetWeather(location string) *dtos.WeatherDto {
	result, err := w.weatherService.GetWeather(location)
	if err != nil {
		log.Err(err).Msg("Failed to get weather")
		return nil
	}
	return result
}
