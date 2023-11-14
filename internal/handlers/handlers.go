package handlers

import (
	"github.com/go-chi/chi"
	chiniddle "github.com/go-chi/chi/middleware"
	"github.com/wenchnoob/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chiniddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}