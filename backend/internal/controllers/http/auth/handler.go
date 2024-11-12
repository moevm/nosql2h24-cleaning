package auth

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
)

type AuthService interface {
	Register(ctx context.Context, user *models.User) (string, error)
	Login(ctx context.Context, user *models.User) (*models.User, *models.Token, error)
	Logout(ctx context.Context, refreshToken string) error
}

type Handler struct {
	auth AuthService
}

func New(auth AuthService) *Handler {
	return &Handler{auth: auth}
}

func (h *Handler) Routes() chi.Router {
	router := chi.NewRouter()
	router.Post("/register", h.Register)
	router.Post("/login", h.Login)
	return router
}
