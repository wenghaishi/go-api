package handlers

import (
	""
	"githib.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/wenghaishi/go-api/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorisation)
		router.Get("/coins", GetCoinBalance)
	})
}
