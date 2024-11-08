package services

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// Create new service
// @Summary      Creates new service
// @Description  Creates new service & returns it's id
// @Tags         Services
// @Accept       json
// @Produce      json
// @Param service body models.Service true "Service info"
// @Success      200  {object}  types.CreateResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/services [post]
func (h *Hander) PostService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	if err := render.DecodeJSON(r.Body, &service); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}
	id, err := h.service.CreateService(r.Context(), &service)
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	render.JSON(w, r, types.CreateResponse{ID: id})
}
