package services

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// Delete service
// @Summary      Deletes service
// @Description  Delete service from database
// @Tags         Services
// @Produce      json
// @Param id path string true "Service id"
// @Success      200
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/services/{id} [delete]
func (h *Hander) DeleteService(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.service.DeleteService(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrServiceNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
	}
}
