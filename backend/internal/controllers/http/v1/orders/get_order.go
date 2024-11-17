package orders

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// GetOrder
// @Summary Get order
// @Description Get order by id
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/orders/{id} [get]
func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {

	orders, err := h.service.GetOrderById(r.Context(), chi.URLParam(r, "id"))
	if err != nil {
		if errors.Is(err, services.ErrOrderNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}

	render.JSON(w, r, orders)
}
