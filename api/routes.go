package api

import (
	"net/http"

	"github.com/fabruun/go-authentication/authentication"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/authentication", loadAuthenticationRoutes)

	return router
}

func loadAuthenticationRoutes(router chi.Router) {
	service := authentication.Authentication{}

	router.Post("/register", service.Register)
}
