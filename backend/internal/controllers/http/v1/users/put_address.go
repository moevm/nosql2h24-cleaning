package users

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Update user address by id
// @Summary      Update user address by id
// @Description  Update user address by id from database
// @Tags         Users
// @Produce      json
// @Param id path string true "user_id"
// @Param address_id path string true "address_id"
// @Param address body models.Address true "address_id"
// @Success      200
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/users/{id}/addresses/{address_id} [put]
func (h *Handler) PutAddress(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	addressID := chi.URLParam(r, "address_id")

	_id, err := bson.ObjectIDFromHex(addressID)
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	var address models.Address
	if err := render.DecodeJSON(r.Body, &address); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}
	address.ID = _id

	if err := h.addressService.UpdateAddress(r.Context(), userID, &address); err != nil {
		if errors.Is(err, services.ErrUserNotFound) || errors.Is(err, services.ErrAddressNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
}
