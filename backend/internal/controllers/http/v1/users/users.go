package users

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/validate"
)

// Get users
// @Summary      Get list of users
// @Description  Get list of users from database
// @Tags         Users
// @Produce      json
// @Param type query string false "specifying user_type"
// @Param name query string false "user name"
// @Param surname query string false "user surname"
// @Param patronymic query string false "user patronymic"
// @Param email query string false "user email"
// @Param phone_number query string false "user phone number"
// @Param orders_count_min query int false "worker orders count min"
// @Param orders_count_max query int false "worker orders count max"
// @Param created_at_begin query string false "created_at begin (string RFC3339)"
// @Param created_at_end query string false "created_at end (string RFC3339)"
// @Success      200  {array}  models.User
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/users [get]
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	filters := types.UserFilters{
		UserData: types.UserData{
			Name:        r.URL.Query().Get("name"),
			Surname:     r.URL.Query().Get("surname"),
			Patronymic:  r.URL.Query().Get("patronymic"),
			Email:       r.URL.Query().Get("email"),
			PhoneNumber: r.URL.Query().Get("phone_number"),
		},
		UserType: r.URL.Query().Get("type"),
	}

	if createdAtBegin, err := time.Parse(time.RFC3339, r.URL.Query().Get("created_at_begin")); err == nil {
		filters.CreatedAt.Begin = &createdAtBegin
	}
	if createdAtEnd, err := time.Parse(time.RFC3339, r.URL.Query().Get("created_at_end")); err == nil {
		filters.CreatedAt.End = &createdAtEnd
	}

	if ordersCountMin, err := strconv.Atoi(r.URL.Query().Get("orders_count_min")); err == nil {
		filters.OrdersCount.Min = &ordersCountMin
	}
	if ordersCountMax, err := strconv.Atoi(r.URL.Query().Get("orders_count_max")); err == nil {
		filters.OrdersCount.Max = &ordersCountMax
	}

	users, err := h.userService.GetUsers(r.Context(), filters)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	render.JSON(w, r, users)
}

// Create user
// @Summary      Creates user with user_type=WORKER
// @Description  Creates user with user_type=WORKER and returns it's id
// @Tags         Users
// @Produce      json
// @Param user body models.User true "New worker info"
// @Success      201  {object}  types.CreateResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/users [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := render.DecodeJSON(r.Body, &user); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}
	user.UserType = "WORKER"

	if err := validate.Validate.Struct(user); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	id, err := h.userService.CreateUser(r.Context(), &user)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, types.CreateResponse{
		ID: id,
	})
}
