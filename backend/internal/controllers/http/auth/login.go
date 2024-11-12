package auth

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/middlewares"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/validate"
)

// Login
// @Summary      Set cookie with jwt
// @Description  If user exist in database set jwt cookie
// @Tags         Auth
// @Accept       json
// @Param		 Creds body models.UserCredentials true "User credentials"
// @Success      200  {object} models.User
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var creds models.UserCredentials
	if err := render.DecodeJSON(r.Body, &creds); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}
	if err := validate.Validate.Struct(creds); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}
	user, token, err := h.auth.Login(r.Context(), &models.User{
		UserCredentials: creds,
	})
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     middlewares.JWTCookie,
		Value:    token.Access,
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60, // 7 days
		HttpOnly: true,
	})
	render.JSON(w, r, user)
}
