package persistence

import "codeberg.org/Birkenfunk/SQS/dtos"

type IDatabase interface {
	AddWeather(dto *dtos.WeatherDto) error
	GetWeather() ([]dtos.WeatherDto, error)
	GetWeatherByLocation(location string) (*dtos.WeatherDto, error)
}
