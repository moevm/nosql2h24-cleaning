package orders

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

// UpdateOrder
// @Summary Update order
// @Description Update order by id
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body models.Order true "Order"
// @Success 200
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/orders/{id} [put]
func (h *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	var order models.Order
	if err := render.DecodeJSON(r.Body, &order); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}
	order.ID = _id

	if err := h.service.UpdateOrder(r.Context(), &order); err != nil {
		if errors.Is(err, services.ErrOrderNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
}
