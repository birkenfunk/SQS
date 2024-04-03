package persistence

import "codeberg.org/Birkenfunk/SQS/dtos"

type DatabaseMock struct {
	weatherDto *dtos.WeatherDto
	err        error
}

func (dm *DatabaseMock) AddWeather(dto *dtos.WeatherDto) error {
	dm.weatherDto = dto
	return dm.err
}

func (dm *DatabaseMock) GetWeather() ([]dtos.WeatherDto, error) {
	return nil, nil
}

func (dm *DatabaseMock) GetWeatherByLocation(_ string) (*dtos.WeatherDto, error) {
	return dm.weatherDto, dm.err
}

func NewDatabaseMock(weatherDto *dtos.WeatherDto, err error) *DatabaseMock {
	return &DatabaseMock{
		weatherDto: weatherDto,
		err:        err,
	}
}
