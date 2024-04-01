package handler

import (
	"codeberg.org/Birkenfunk/SQS/business/logic"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

var weather logic.IWeather = &logic.Weather{}

// WeatherHandler is a handler for the Weather Endpoint.
func WeatherHandler(rw http.ResponseWriter, r *http.Request) {
	location := chi.URLParam(r, "location")
	weather := weather.GetWeather(location)
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
