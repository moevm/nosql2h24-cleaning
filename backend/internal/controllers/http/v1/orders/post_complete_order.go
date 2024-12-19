package orders

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// CompleteOrder
// @Summary      Worker complete order
// @Description  Worker complete order
// @Tags         Orders
// @Accept      json
// @Produce      json
// @Param id path string true "Order id"
// @Success      200
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/orders/{id}/complete [post]
func (h *Handler) CompleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id")

	if err := h.service.CompleteOrder(r.Context(), orderID); err != nil {
		if errors.Is(err, services.ErrOrderAlreadyDone) {
			render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
}
