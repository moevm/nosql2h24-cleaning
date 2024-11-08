package users

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// Get user addresses
// @Summary      Get list of user addresses
// @Description  Get list of user addresses from database
// @Tags         Users
// @Produce      json
// @Param id path string true "user_id"
// @Success      200  {array}  models.Address
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/users/{id}/addresses [get]
func (h *Handler) GetAddresses(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	addresses, err := h.addressService.GetAddresses(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	// print not null
	if addresses == nil {
		addresses = make([]*models.Address, 0)
	}
	render.JSON(w, r, addresses)
}
