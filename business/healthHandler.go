package business

import "net/http"

// HealthHandler is a handler for health check endpoint.
func HealthHandler(rw http.ResponseWriter, _ *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("OK"))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
