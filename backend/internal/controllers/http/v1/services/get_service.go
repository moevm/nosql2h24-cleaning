package services

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// Get service
// @Summary      Get service info
// @Description  Get service from database
// @Tags         Services
// @Produce      json
// @Param id path string true "Service id"
// @Success      200  {object}  models.Service
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/services/{id} [get]
func (h *Hander) GetService(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	service, err := h.service.GetService(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrServiceNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
	}
	render.JSON(w, r, service)
}
