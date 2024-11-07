package auth

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// Register
// @Summary      Create user with user_type=CLIENT
// @Description  Create user with user_type=CLIENT in database
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param user body models.NewUser true "New user info"
// @Success      200  {object}  types.CreateResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/auth/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := render.DecodeJSON(r.Body, &user); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}
	user.UserType = "CLIENT"

	id, err := h.auth.Register(r.Context(), &user)
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	render.JSON(w, r, types.CreateResponse{
		ID: id,
	})
}
