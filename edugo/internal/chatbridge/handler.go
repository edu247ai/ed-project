package chatbridge

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	service *Service
}

func NewHandler(pool *pgxpool.Pool) *Handler {
	repo := NewRepository(pool)
	ws := NewWebSocketBridge()
	service := NewService(repo, ws)
	return &Handler{service: service}
}

type sendMessageRequest struct {
	Content string `json:"content"`
}

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	roomID, _ := strconv.Atoi(ps.ByName("room_id"))
	userID := 1 // TODO: replace with real auth

	var req sendMessageRequest
	json.NewDecoder(r.Body).Decode(&req)

	msg, err := h.service.SendMessage(ctx, roomID, userID, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(msg)
}

func (h *Handler) ListMessages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	roomID, _ := strconv.Atoi(ps.ByName("room_id"))

	msgs, err := h.service.GetRoomMessages(ctx, roomID)
	if err != nil {
		http.Error(w, "failed to load messages", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(msgs)
}
