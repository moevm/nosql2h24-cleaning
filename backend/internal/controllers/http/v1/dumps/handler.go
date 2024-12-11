package dumps

import (
	"github.com/go-chi/chi/v5"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
)

type Handler struct {
	service *services.DumpService
}

func NewHandler(service *services.DumpService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", h.GetDump)
	router.Post("/", h.PostDump)

	return router
}
