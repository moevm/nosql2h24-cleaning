package users

import (
	"github.com/go-chi/chi/v5"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
)

type Handler struct {
	service *services.UserService
}

func New(service *services.UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", h.GetUsers)
	router.Post("/", h.CreateUser)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.GetUser)
		r.Put("/", h.UpdateUser)
		r.Delete("/", h.DeleteUser)
	})
	return router
}
