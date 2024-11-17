package orders

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// GetOrders
// @Summary Get orders
// @Description Get all orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param user_id query string false "User ID"
// @Param workers_id query []string false "Workers ID"
// @Param statuses query []string false "Statuses"
// @Success 200 {array} models.Order
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/orders [get]
func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID := query.Get("user_id")
	workersID := query["workers_id"]
	statuses := query["statuses"]
	orders, err := h.service.GetOrders(r.Context(), types.SearchParams{
		UserID:    userID,
		WorkersID: workersID,
		Statuses:  statuses,
	})
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}

	render.JSON(w, r, orders)
}
