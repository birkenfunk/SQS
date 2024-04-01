package service

import (
	"codeberg.org/Birkenfunk/SQS/dtos"
)

type WeatherServiceMock struct {
	Error   error
	Weather *dtos.WeatherDto
}

func (ws *WeatherServiceMock) GetWeather(_ string) (*dtos.WeatherDto, error) {
	return ws.Weather, ws.Error
}

func (ws *WeatherServiceMock) GetHealth() error {
	return ws.Error
}
