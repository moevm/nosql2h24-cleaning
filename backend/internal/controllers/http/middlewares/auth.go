package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

type ContextKey string

const (
	JWTCookie            = "token"
	UserID    ContextKey = "user-id"
)

type JWT interface {
	Parse(token string) (string, error)
}

type AuthMiddleware struct {
	jwt JWT
}

func NewAuthMiddleware(jwt JWT) *AuthMiddleware {
	return &AuthMiddleware{
		jwt: jwt,
	}
}

func (m *AuthMiddleware) JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(JWTCookie)
		if err != nil || len(cookie.Value) == 0 {
			render.Render(w, r, httputil.NewError(http.StatusUnauthorized, err))
			return
		}

		id, err := m.jwt.Parse(cookie.Value)
		if err != nil {
			render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
			return
		}
		ctx := context.WithValue(r.Context(), UserID, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
