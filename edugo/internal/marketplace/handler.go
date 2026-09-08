package marketplace

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"edu247/edugo/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	db      *pgxpool.Pool
	storage *storage.Manager
	service *Service
}

// NewHandler creates a new marketplace HTTP handler.
func NewHandler(pool *pgxpool.Pool, store *storage.Manager) *Handler {
	repo := NewRepository(pool)
	service := NewService(repo)

	return &Handler{
		db:      pool,
		storage: store,
		service: service,
	}
}

//
// ------------------------------------------------------------
// GET /api/products
// ------------------------------------------------------------
//

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	products, err := h.service.ListActiveProducts(ctx)
	if err != nil {
		http.Error(w, "failed to list products", http.StatusInternalServerError)
		return
	}

	writeJSON(w, products)
}

//
// ------------------------------------------------------------
// GET /api/products/:id
// ------------------------------------------------------------
//

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	idStr := ps.ByName("id")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProduct(ctx, productID)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	writeJSON(w, product)
}

//
// ------------------------------------------------------------
// POST /api/orders
// ------------------------------------------------------------
//

type createOrderRequest struct {
	Items []struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	} `json:"items"`
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	// TODO: Replace with real auth middleware
	userID := 1

	var req createOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	var items []OrderItemInput
	for _, it := range req.Items {
		items = append(items, OrderItemInput{
			ProductID: it.ProductID,
			Quantity:  it.Quantity,
		})
	}

	order, err := h.service.CreateOrder(ctx, userID, items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, order)
}

//
// ------------------------------------------------------------
// GET /api/files/:id/download
// ------------------------------------------------------------
//

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	fileID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	// TODO: Replace with real auth middleware
	userID := 1

	allowed, err := h.service.CanUserAccessFile(ctx, userID, fileID)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	if !allowed {
		http.Error(w, "access denied", http.StatusForbidden)
		return
	}

	// Load file metadata
	file, err := h.service.repo.GetFileByID(ctx, fileID)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	// Generate signed MinIO URL (valid for 15 minutes)
	url, err := h.storage.PresignURL(file.StoragePath, 15*time.Minute)
	if err != nil {
		http.Error(w, "cannot generate download link", http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{
		"download_url": url,
		"expires_in":   "900",
	})
}

//
// ------------------------------------------------------------
// Helper: JSON writer
// ------------------------------------------------------------
//

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
