package services

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

// Update service
// @Summary      Update service info
// @Description  Update service info in database, if exist
// @Tags         Services
// @Accept       json
// @Produce      json
// @Param id path string true "Service id"
// @Param service body models.Service true "Service info"
// @Success      200
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/services/{id} [put]
func (h *Hander) PutService(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	var service models.Service
	if err := render.DecodeJSON(r.Body, &service); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}
	service.ID = _id

	if err := h.service.UpdateService(r.Context(), &service); err != nil {
		if errors.Is(err, services.ErrServiceNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
}
