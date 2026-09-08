package files

import (
	"encoding/json"
	"net/http"

	"edu247/edugo/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	service *Service
}

func NewHandler(pool *pgxpool.Pool, manager *storage.Manager) *Handler {
	return &Handler{
		service: NewService(pool, manager),
	}
}

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	// TODO: replace with JWT auth
	userID := 1

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "missing file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	saved, err := h.service.SaveFile(ctx, userID, file, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(saved)
}
