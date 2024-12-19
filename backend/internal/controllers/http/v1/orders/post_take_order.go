package orders

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/middlewares"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// TakeOrder
// @Summary      Worker take order
// @Description  Worker take order
// @Tags         Orders
// @Accept      json
// @Produce      json
// @Param id path string true "Order id"
// @Success      200
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/orders/{id}/take [post]
func (h *Handler) TakeOrder(w http.ResponseWriter, r *http.Request) {
	workerID := r.Context().Value(middlewares.UserID).(string)
	orderID := chi.URLParam(r, "id")

	if err := h.service.AssignWorkerToOrder(r.Context(), orderID, workerID); err != nil {
		if errors.Is(err, services.ErrOrderAlreadyDone) {
			render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
}
