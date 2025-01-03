package users

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// Get user address by id
// @Summary      Get user address
// @Description  Get user address from database
// @Tags         Users
// @Produce      json
// @Param id path string true "user_id"
// @Param address_id path string true "address_id"
// @Success      200  {object}  models.AddressWithTimestamp
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/users/{id}/addresses/{address_id} [get]
func (h *Handler) GetAddress(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	addressID := chi.URLParam(r, "address_id")

	address, err := h.addressService.GetAddress(r.Context(), userID, addressID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) || errors.Is(err, services.ErrAddressNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	render.JSON(w, r, address)
}
