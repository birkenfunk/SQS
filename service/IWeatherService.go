package service

type IWeatherService interface {
	GetWeather(location string) (string, error)
	GetHealth() (string, error)
}
