package auth

import (
	"net/http"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/middlewares"
)

// Logout
// @Summary      Remove cookie with jwt
// @Description  If cookie exist, remove it
// @Tags         Auth
// @Success      200
// @Failure      400  {object}  httputil.HTTPError
// @Router       /api/auth/logout [post]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     middlewares.JWTCookie,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}
