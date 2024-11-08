package users

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// Create new user address
// @Summary      Create new user address
// @Description  Create new user address from database
// @Tags         Users
// @Produce      json
// @Param id path string true "user_id"
// @Param address body models.Address true "address info"
// @Success      200  {object}   types.CreateResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/users/{id}/addresses [post]
func (h *Handler) PostAddress(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var address models.Address
	if err := render.DecodeJSON(r.Body, &address); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	id, err := h.addressService.CreateAddress(r.Context(), id, &address)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	render.JSON(w, r, types.CreateResponse{ID: id})
}
