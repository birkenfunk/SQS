package business

import "net/http"

// HealthHandler is a handler for health check endpoint.
func WeatherHandler(rw http.ResponseWriter, _ *http.Request) {
	rw.WriteHeader(http.StatusNotImplemented)
	_, err := rw.Write([]byte("Not implemented yet"))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
