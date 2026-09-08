package http

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"

	"edu247/edugo/internal/ai"
	"edu247/edugo/internal/auth"
	"edu247/edugo/internal/chatbridge"
	"edu247/edugo/internal/config"
	"edu247/edugo/internal/courses"
	"edu247/edugo/internal/files"
	"edu247/edugo/internal/marketplace"
	"edu247/edugo/internal/storage"
	"edu247/edugo/internal/users"
)

// JSON helper
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func NewRouter(pool *pgxpool.Pool, cfg config.Config, store *storage.Manager) http.Handler {
	r := httprouter.New()

	//
	// USERS
	//
	userHandler := users.NewHandler(pool)
	r.POST("/api/users/register", userHandler.Register)
	r.POST("/api/users/login", userHandler.Login)
	r.GET("/api/users/:id", userHandler.GetProfile)

	//
	// AUTH
	//
	authHandler := auth.NewHandler(pool, cfg)
	r.POST("/api/auth/register", authHandler.Register)
	r.POST("/api/auth/login", authHandler.Login)

	//
	// COURSES
	//
	courseHandler := courses.NewHandler(pool)
	r.POST("/api/courses", courseHandler.CreateCourse)
	r.GET("/api/courses", courseHandler.ListCourses)
	r.GET("/api/courses/:id", courseHandler.GetCourse)
	r.POST("/api/courses/:id/lessons", courseHandler.AddLesson)
	r.GET("/api/courses/:id/lessons", courseHandler.ListLessons)

	//
	// FILES
	//
	fileHandler := files.NewHandler(pool, store)
	r.POST("/api/files/upload", fileHandler.Upload)

	//
	// MARKETPLACE
	//
	marketHandler := marketplace.NewHandler(pool, store)
	r.GET("/api/products", marketHandler.ListProducts)
	r.GET("/api/products/:id", marketHandler.GetProduct)
	r.POST("/api/orders", marketHandler.CreateOrder)
	r.GET("/api/files/:id/download", marketHandler.DownloadFile)

	//
	// CHATBRIDGE
	//
	chatHandler := chatbridge.NewHandler(pool)
	r.POST("/api/chat/:room_id/send", chatHandler.SendMessage)
	r.GET("/api/chat/:room_id/messages", chatHandler.ListMessages)

	//
	// ------------------------------------------------------------
	// AI ROUTES (EDU247MODEL Integration)
	// ------------------------------------------------------------
	//

	// RAG Query
	r.GET("/api/ai/rag", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		query := req.URL.Query().Get("q")
		if query == "" {
			writeJSON(w, 400, map[string]string{"error": "missing q"})
			return
		}

		result, err := ai.RAGQuery(query)
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, 200, result)
	})

	// Anomaly Detection
	r.GET("/api/ai/anomaly", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		result, err := ai.Anomaly(0.78)
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, result)
	})

	// Forecast
	r.POST("/api/ai/forecast", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		var body struct {
			Series []float64 `json:"series"`
			Steps  int       `json:"steps"`
		}

		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			writeJSON(w, 400, map[string]string{"error": "invalid JSON"})
			return
		}

		result, err := ai.Forecast(body.Series, body.Steps)
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, 200, result)
	})

	// RL Evaluation
	r.GET("/api/ai/rl", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		result, err := ai.RL()
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, result)
	})

	// Skills List
	r.GET("/api/ai/skills", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		skills, err := ai.Skills()
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, map[string]interface{}{"skills": skills})
	})

	//
	// MIDDLEWARE WRAP
	//
	final := CORSMiddleware(LoggingMiddleware(r))
	return final
}
