package http

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

//
// ------------------------------------------------------------
// JSON Helpers
// ------------------------------------------------------------
//

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, map[string]string{"error": msg})
}

//
// ------------------------------------------------------------
// Request ID Middleware
// ------------------------------------------------------------
//

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		r.Header.Set("X-Request-ID", id)
		next.ServeHTTP(w, r)
	})
}

//
// ------------------------------------------------------------
// Logging Middleware
// ------------------------------------------------------------
//

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s (%s)", r.Header.Get("X-Request-ID"), r.Method, r.URL.Path, time.Since(start))
	})
}

//
// ------------------------------------------------------------
// Panic Recovery Middleware
// ------------------------------------------------------------
//

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("PANIC: %v\n%s", rec, debug.Stack())
				WriteError(w, http.StatusInternalServerError, "internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

//
// ------------------------------------------------------------
// CORS Middleware
// ------------------------------------------------------------
//

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

//
// ------------------------------------------------------------
// Security Headers Middleware
// ------------------------------------------------------------
//

func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}

//
// ------------------------------------------------------------
// Rate Limiting Middleware (simple global limiter)
// ------------------------------------------------------------
//

var (
	rateLimit      = 50
	rateWindow     = time.Second
	requestCounter = 0
	mu             sync.Mutex
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		requestCounter++
		count := requestCounter
		mu.Unlock()

		if count > rateLimit {
			WriteError(w, http.StatusTooManyRequests, "rate limit exceeded")
			return
		}

		next.ServeHTTP(w, r)

		// Reset counter every second
		go func() {
			time.Sleep(rateWindow)
			mu.Lock()
			requestCounter = 0
			mu.Unlock()
		}()
	})
}

//
// ------------------------------------------------------------
// JWT Extraction Middleware (token only)
// ------------------------------------------------------------
//

func GetAuthToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return ""
}

//
// ------------------------------------------------------------
// Role-Based Access Middleware (teacher/admin/student)
// ------------------------------------------------------------
//

func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role := r.Header.Get("X-User-Role") // set by your auth handler

			for _, allowed := range allowedRoles {
				if role == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}

			WriteError(w, http.StatusForbidden, "insufficient permissions")
		})
	}
}
