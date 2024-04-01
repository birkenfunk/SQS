package presentation

import (
	"codeberg.org/Birkenfunk/SQS/business/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()

	// Add Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.RealIP)
	r.Use(middleware.RedirectSlashes)

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

	r.Get("/health", handler.HealthHandler)
	r.Get("/weather/{location}", handler.WeatherHandler)
	return r
}
