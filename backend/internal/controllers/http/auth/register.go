package auth

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/types"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/validate"
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
	var creds models.UserCredentials
	if err := render.DecodeJSON(r.Body, &creds); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	if err := validate.Validate.Struct(creds); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	id, err := h.auth.Register(r.Context(), &models.User{
		UserCredentials: creds,
		UserType:        "CLIENT",
	})
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
	render.JSON(w, r, types.CreateResponse{
		ID: id,
	})
}
