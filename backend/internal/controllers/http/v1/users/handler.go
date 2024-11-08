package users

import (
	"github.com/go-chi/chi/v5"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
)

type Handler struct {
	userService    *services.UserService
	addressService *services.AddressService
}

func New(
	userService *services.UserService,
	addressService *services.AddressService,
) *Handler {
	return &Handler{
		userService:    userService,
		addressService: addressService,
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

		r.Get("/addresses", h.GetAddresses)
		r.Post("/addresses", h.PostAddress)

		r.Route("/addresses/{address_id}", func(r chi.Router) {
			r.Get("/", h.GetAddress)
			r.Put("/", h.PutAddress)
			r.Delete("/", h.DeleteAddress)
		})
	})

	return router
}
