package services

import (
	"github.com/go-chi/chi/v5"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
)

type Hander struct {
	service *services.Service
}

func New(service *services.Service) *Hander {
	return &Hander{
		service: service,
	}
}

func (h *Hander) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", h.GetServices)
	router.Post("/", h.PostService)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.GetService)
		r.Put("/", h.PutService)
		r.Delete("/", h.DeleteService)
	})

	return router
}
