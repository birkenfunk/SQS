package handler

import "net/http"

type IHealthHandler interface {
	GetHealthHandler(rw http.ResponseWriter, r *http.Request)
}

type HealthHandler struct{}

func NewHealthHandler() IHealthHandler {
	return &HealthHandler{}
}

// GetHealthHandler is a handler for health check endpoint.
func (wh *HealthHandler) GetHealthHandler(rw http.ResponseWriter, _ *http.Request) {
	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("OK"))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
