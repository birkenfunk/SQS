package service

import "codeberg.org/Birkenfunk/SQS/dtos"

type IWeatherService interface {
	GetWeather(location string) (*dtos.WeatherDto, error)
	GetHealth() error
}
