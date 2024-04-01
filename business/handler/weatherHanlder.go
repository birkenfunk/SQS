package business

import (
	"encoding/json"
	"net/http"
)

var weather IWeather = &Weather{}

// WeatherHandler is a handler for the Weather Endpoint.
func WeatherHandler(rw http.ResponseWriter, _ *http.Request) {
	weather := weather.getWeather("Berlin")
	if weather == nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, err := rw.Write([]byte("Failed to get weather"))
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	response, err := json.Marshal(weather)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(response)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
