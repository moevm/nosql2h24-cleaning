package dumps

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httputil"
)

// Get dump
// @Summary      Get dump as json file
// @Description  Get dump as json file
// @Tags         Dump
// @Produce      octet-stream
// @Success      200  {file} file
// @Failure      500  {object}  httputil.HTTPError
// @Router       /api/v1/dumps/ [get]
func (h *Handler) GetDump(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.Export(r.Context())
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}

	file, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		render.Render(w, r, httputil.NewError(http.StatusInternalServerError, err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename="+"dump.json")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Write(file)
}
