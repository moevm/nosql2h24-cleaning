package orders

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/middlewares"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// GetOrders
// @Summary Get orders
// @Description Get orders from database. If user_id is "me" then it will be replaced with the current user ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param user_id query string false "User ID"
// @Param user_name query string false "User name"
// @Param user_surname query string false "User surname"
// @Param user_patronymic query string false "User patronymic"
// @Param user_email query string false "User email"
// @Param address_city query string false "Address city"
// @Param address_street query string false "Address street"
// @Param address_building query string false "Address building"
// @Param address_entrance query string false "Address entrance"
// @Param address_floor query string false "Address floor"
// @Param address_door_number query string false "Address door number"
// @Param price_min query int false "Price min"
// @Param price_max query int false "Price max"
// @Param area_min query int false "Area min"
// @Param area_max query int false "Area max"
// @Param number_of_rooms_min query int false "Number of rooms min"
// @Param number_of_rooms_max query int false "Number of rooms max"
// @Param number_of_bathrooms_min query int false "Number of bathrooms min"
// @Param number_of_bathrooms_max query int false "Number of bathrooms max"
// @Param pollution_min query int false "Pollution min"
// @Param pollution_max query int false "Pollution max"
// @Param required_workers_min query int false "Required workers min"
// @Param required_workers_max query int false "Required workers max"
// @Param worker_name query string false "Worker name"
// @Param worker_surname query string false "Worker surname"
// @Param worker_patronymic query string false "Worker patronymic"
// @Param worker_email query string false "Worker email"
// @Param workers_id query []string false "Workers ID"
// @Param statuses query []string false "Statuses"
// @Param services query []string false "Services"
// @Param date_time_begin query string false "Date time begin (string RFC3339)"
// @Param date_time_end query string false "Date time end (string RFC3339)"
// @Success 200 {array} models.Order
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/orders [get]
func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filters := types.OrderFilters{
		UserID: query.Get("user_id"),
		User: types.UserData{
			Name:       query.Get("user_name"),
			Surname:    query.Get("user_surname"),
			Patronymic: query.Get("user_patronymic"),
			Email:      query.Get("user_email"),
		},
		Worker: types.UserData{
			Name:       query.Get("worker_name"),
			Surname:    query.Get("worker_surname"),
			Patronymic: query.Get("worker_patronymic"),
			Email:      query.Get("worker_email"),
		},
		Address: types.AddressData{
			City:       query.Get("address_city"),
			Street:     query.Get("address_street"),
			Building:   query.Get("address_building"),
			Entrance:   query.Get("address_entrance"),
			Floor:      query.Get("address_floor"),
			DoorNumber: query.Get("address_door_number"),
		},
		Price:             types.NewNumberRange("price", query),
		Area:              types.NewNumberRange("area", query),
		NumberOfRooms:     types.NewNumberRange("number_of_rooms", query),
		NumberOfBathrooms: types.NewNumberRange("number_of_bathrooms", query),
		Pollution:         types.NewNumberRange("pollution", query),
		RequiredWorkers:   types.NewNumberRange("required_workers", query),
		WorkersID:         query["workers_id"],
		Statuses:          query["statuses"],
		Services:          query["services"],
		DateTime:          types.NewTimeInterval("date_time", query),
	}

	if filters.UserID == "me" {
		filters.UserID = middlewares.GetUserID(r.Context())
	}

	orders, err := h.service.GetOrders(r.Context(), filters)
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	if orders == nil {
		orders = []*models.Order{}
	}
	render.JSON(w, r, orders)
}
