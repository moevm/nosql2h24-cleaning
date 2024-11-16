package users

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/middlewares"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/validate"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Get user
// @Summary      Get user info
// @Description  Get user from database
// @Tags         Users
// @Produce      json
// @Param id path string true "User id"
// @Success      200  {object}  models.User
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/users/{id} [get]
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "me" {
		id = r.Context().Value(middlewares.UserID).(string)
	}
	user, err := h.userService.GetUserById(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	render.JSON(w, r, user)
}

// Update user
// @Summary      Update user
// @Description  Update user info
// @Tags         Users
// @Accept      json
// @Produce      json
// @Param id path string true "User id"
// @Param user body models.User true "User data"
// @Success      200
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/users/{id} [put]
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := chi.URLParam(r, "id")

	_id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	if err := render.DecodeJSON(r.Body, &user); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}
	user.ID = _id

	if err := validate.Validate.Struct(user); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	if err := h.userService.UpdateUser(r.Context(), &user); err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
}

// Delete user
// @Summary      Delete user
// @Description  Delete user from database
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param id path string true "User id"
// @Success      204
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/users/{id} [delete]
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if _, err := bson.ObjectIDFromHex(id); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	if err := h.userService.DeleteUser(r.Context(), id); err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			render.Render(w, r, httputil.NewError(http.StatusNotFound, err))
			return
		}
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	render.Status(r, http.StatusNoContent)
}
