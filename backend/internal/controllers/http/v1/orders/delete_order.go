package orders

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// DeleteOrder
// @Summary Delete order
// @Description Delete order by id
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/orders/{id} [delete]
func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id")
	if _, err := bson.ObjectIDFromHex(orderID); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	if err := h.service.DeleteOrder(r.Context(), orderID); err != nil {
		if errors.Is(err, services.ErrOrderNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
