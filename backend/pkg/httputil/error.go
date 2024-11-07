package httputil

import (
	"net/http"

	"github.com/go-chi/render"
)

func NewError(status int, err error) render.Renderer {
	return &HTTPError{
		Error:          err,
		HTTPStatusCode: status,
		StatusText:     http.StatusText(status),
		ErrorText:      err.Error(),
	}
}

type HTTPError struct {
	Error          error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status" example:"bad request"`
	ErrorText      string `json:"error" example:"password field is empty"`
}

func (h *HTTPError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, h.HTTPStatusCode)
	return nil
}
