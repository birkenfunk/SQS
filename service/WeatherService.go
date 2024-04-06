package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"codeberg.org/Birkenfunk/SQS/consts"
	"codeberg.org/Birkenfunk/SQS/dtos"
	"github.com/rs/zerolog/log"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WeatherService struct {
	client HTTPClient
}

func NewWeatherService() IWeatherService {
	return &WeatherService{client: http.DefaultClient}
}

var errHealthEndpointAvailable = errors.New("health endpoint is not available")

func (ws *WeatherService) GetHealth() error {
	request, err := http.NewRequest("GET", consts.GetWeatherServiceURL()+"/api/v1/health", nil)
	if err != nil {
		log.Error().Msgf("client: could not create request: %s\n", err)
		return err
	}

	res, err := ws.client.Do(request)
	if err != nil {
		log.Error().Msgf("client: could not send request: %s\n", err)
		return err
	}
	if res.StatusCode != 200 {
		log.Error().Msgf("client: health endpoint is not available: %d\n", res.StatusCode)
		return fmt.Errorf("client: health endpoint is not available: %w", errHealthEndpointAvailable)
	}
	err = res.Body.Close()
	if err != nil {
		log.Error().Msgf("client: could not close response body: %s\n", err)
		return err
	}
	return nil
}

func (ws *WeatherService) GetWeather(location string) (*dtos.WeatherDto, error) {
	var weatherDto dtos.WeatherDto
	request, err := http.NewRequest("GET", consts.GetWeatherServiceURL()+"/api/v1/weather/"+location, nil)
	if err != nil {
		log.Error().Msgf("client: could not create request: %s\n", err)
		return nil, err
	}

	res, err := ws.client.Do(request)
	if err != nil {
		log.Error().Msgf("client: could not send request: %s\n", err)
		return nil, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Msgf("client: could not read response: %s\n", err)
		return nil, err
	}
	err = res.Body.Close()
	if err != nil {
		log.Error().Msgf("client: could not close response body: %s\n", err)
		return nil, err
	}
	err = json.Unmarshal(resBody, &weatherDto)
	if err != nil {
		log.Error().Msgf("client: could not unmarshal response: %s\n", err)
		return nil, err
	}
	return &weatherDto, nil
}
