package handler

import (
	"encoding/json"
	"net/http"

	"codeberg.org/Birkenfunk/SQS/business/logic"
	"github.com/go-chi/chi/v5"
)

type IWeatherHandler interface {
	GetWeatherHandler(rw http.ResponseWriter, r *http.Request)
}

type WeatherHandler struct {
	weather logic.IWeather
}

func NewWeatherHandler() IWeatherHandler {
	return &WeatherHandler{weather: logic.NewWeather()}
}

// GetWeatherHandler is a handler for the Weather Endpoint.
func (h *WeatherHandler) GetWeatherHandler(rw http.ResponseWriter, r *http.Request) {
	location := chi.URLParam(r, "location")
	weather := h.weather.GetWeather(location)
	if weather == nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, err := rw.Write([]byte("Failed to get weather"))
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	response, err := json.MarshalIndent(weather, "", "  ")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(response)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// SetWeather is a setter for the weather logic.
func (h *WeatherHandler) SetWeather(w logic.IWeather) {
	h.weather = w
}
