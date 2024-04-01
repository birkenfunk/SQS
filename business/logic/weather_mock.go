package logic

import "codeberg.org/Birkenfunk/SQS/dtos"

type WeatherMock struct {
	WeatherDto *dtos.WeatherDto
}

func (wm *WeatherMock) GetWeather(_ string) *dtos.WeatherDto {
	return wm.WeatherDto
}
