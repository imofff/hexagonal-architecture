package http

import (
	"hexagonal/internal/adapter/http/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(userHandler *handler.UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
		r.Get("/", userHandler.GetAll)
	})
	return r
}
