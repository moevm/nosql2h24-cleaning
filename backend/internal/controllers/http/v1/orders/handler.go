package orders

import (
	"github.com/go-chi/chi/v5"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
)

type Handler struct {
	service *services.OrderService
}

func NewHandler(service *services.OrderService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", h.GetOrders)
	router.Post("/", h.CreateOrder)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.GetOrder)
		r.Put("/", h.UpdateOrder)
		r.Delete("/", h.DeleteOrder)

		r.Post("/complete", h.CompleteOrder)
		r.Post("/take", h.TakeOrder)
	})

	return router
}
