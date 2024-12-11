package dumps

import (
	"io"
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// Post dump
// @Summary      Post dump as json file
// @Description  Post dump as json file
// @Tags         Dump
// @Accept       mpfd
// @Param        file formData file true "json file"
// @Success      200  {file} file
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/dumps/ [post]
func (h *Handler) PostDump(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusBadRequest, err))
		return
	}

	if err := h.service.Import(r.Context(), data, true); err != nil {
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}
}
