package httpserver

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/llevintza/openswim/apps/api/internal/health"
	"github.com/llevintza/openswim/apps/api/internal/metrics"
)

// Server is the API HTTP surface.
type Server struct {
	pool *pgxpool.Pool
	mux  *http.ServeMux
}

// New builds the HTTP mux with health and reserved ops routes.
func New(pool *pgxpool.Pool) *Server {
	s := &Server{
		pool: pool,
		mux:  http.NewServeMux(),
	}
	s.mux.HandleFunc("GET /health", s.handleHealth)
	health.RegisterRoutes(s.mux)
	metrics.RegisterRoutes(s.mux)
	return s
}

// Handler returns the root handler.
func (s *Server) Handler() http.Handler {
	return s.mux
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")
	if err := s.pool.Ping(ctx); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "degraded",
			"error":  "database_unavailable",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
