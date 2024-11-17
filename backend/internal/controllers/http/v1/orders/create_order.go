package orders

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// CreateOrder
// @Summary Create order
// @Description Create order
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order"
// @Success 201 {object} types.CreateResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/orders [post]
func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := render.DecodeJSON(r.Body, &order); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	id, err := h.service.CreateOrder(r.Context(), &order)
	if err != nil {
		if errors.Is(err, services.ErrOrderAlreadyExist) {
			render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, types.CreateResponse{ID: id})
}
