package courses

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
// POST /api/courses
// ------------------------------------------------------------
//

type createCourseRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	// TODO: replace with real auth
	teacherID := 1

	var req createCourseRequest
	json.NewDecoder(r.Body).Decode(&req)

	course, err := h.service.CreateCourse(ctx, teacherID, req.Title, req.Description, req.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, course)
}

//
// ------------------------------------------------------------
// GET /api/courses
// ------------------------------------------------------------
//

func (h *Handler) ListCourses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	courses, err := h.service.ListCourses(ctx)
	if err != nil {
		http.Error(w, "failed to list courses", http.StatusInternalServerError)
		return
	}

	writeJSON(w, courses)
}

//
// ------------------------------------------------------------
// GET /api/courses/:id
// ------------------------------------------------------------
//

func (h *Handler) GetCourse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	id, _ := strconv.Atoi(ps.ByName("id"))
	course, err := h.service.GetCourse(ctx, id)
	if err != nil {
		http.Error(w, "course not found", http.StatusNotFound)
		return
	}

	writeJSON(w, course)
}

//
// ------------------------------------------------------------
// POST /api/courses/:id/lessons
// ------------------------------------------------------------
//

type createLessonRequest struct {
	Title      string `json:"title"`
	FileID     int    `json:"file_id"`
	OrderIndex int    `json:"order_index"`
	Type       string `json:"type"` // pdf, video, audio, quiz
}

func (h *Handler) AddLesson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	courseID, _ := strconv.Atoi(ps.ByName("id"))

	var req createLessonRequest
	json.NewDecoder(r.Body).Decode(&req)

	lesson, err := h.service.AddLesson(
		ctx,
		courseID,
		req.Title,
		req.FileID,
		req.OrderIndex,
		req.Type,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, lesson)
}

//
// ------------------------------------------------------------
// GET /api/courses/:id/lessons
// ------------------------------------------------------------
//

func (h *Handler) ListLessons(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	courseID, _ := strconv.Atoi(ps.ByName("id"))

	lessons, err := h.service.ListLessons(ctx, courseID)
	if err != nil {
		http.Error(w, "failed to list lessons", http.StatusInternalServerError)
		return
	}

	writeJSON(w, lessons)
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
