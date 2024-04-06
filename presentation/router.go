package presentation

import (
	"codeberg.org/Birkenfunk/SQS/business/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type IRouter interface {
	InitRouter() *chi.Mux
}

type Router struct {
	weatherHandler handler.IWeatherHandler
	healthHandler  handler.IHealthHandler
}

func NewRouter() IRouter {
	return &Router{
		weatherHandler: handler.NewWeatherHandler(),
		healthHandler:  handler.NewHealthHandler(),
	}
}

func (router *Router) InitRouter() *chi.Mux {
	r := chi.NewRouter()

	// Add Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.RealIP)
	r.Use(middleware.RedirectSlashes)

	r.Mount("/api", router.apiRouter())

	return r
}

func (router *Router) apiRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/v1", router.v1Router())

	return r
}

func (router *Router) v1Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", router.healthHandler.GetHealthHandler)
	r.Get("/weather/{location}", router.weatherHandler.GetWeatherHandler)
	return r
}
