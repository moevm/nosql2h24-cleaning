package services

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// Get services
// @Summary      Get array of services info
// @Description  Get array of services info from database
// @Tags         Services
// @Produce      json
// @Success      200  {array}  models.Service
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/services [get]
func (h *Hander) GetServices(w http.ResponseWriter, r *http.Request) {
	services, err := h.service.GetServices(r.Context())
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	render.JSON(w, r, services)
}
