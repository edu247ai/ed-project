package users

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
	service := NewService(repo)
	return &Handler{service: service}
}

//
// ------------------------------------------------------------
// POST /api/users/register
// ------------------------------------------------------------
//

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	var req registerRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, err := h.service.Register(ctx, req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, user)
}

//
// ------------------------------------------------------------
// POST /api/users/login
// ------------------------------------------------------------
//

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	var req loginRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, err := h.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	writeJSON(w, user)
}

//
// ------------------------------------------------------------
// GET /api/users/:id
// ------------------------------------------------------------
//

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	id, _ := strconv.Atoi(ps.ByName("id"))

	user, err := h.service.GetProfile(ctx, id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	writeJSON(w, user)
}

//
// ------------------------------------------------------------
// Helper
// ------------------------------------------------------------
//

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
