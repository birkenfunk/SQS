package presentation

import (
	"codeberg.org/Birkenfunk/SQS/business"
	"github.com/go-chi/chi/v5"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/api", apiRouter())

	return r
}

func apiRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/v1", v1Router())

	return r
}

func v1Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", business.HealthHandler)

	return r
}
