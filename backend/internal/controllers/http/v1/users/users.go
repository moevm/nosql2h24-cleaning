package users

import (
	"errors"
	"net/http"

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
// @Param type query string true "specifying user_type"
// @Success      200  {array}  models.User
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/users [get]
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	user_type := r.URL.Query().Get("type")
	if len(user_type) == 0 {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, errors.New("required query param didn't provided")))
		return
	}

	users, err := h.userService.GetUsers(r.Context(), user_type)
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
