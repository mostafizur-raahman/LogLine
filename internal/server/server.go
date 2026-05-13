package server

import (
	"log/slog"
	"logline/internal/config"
	"net/http"
)

type Server struct {
	config config.Config
	logger *slog.Logger
	mux    *http.ServeMux
}

func New(cfg config.Config) *Server {

	logger := newLogger(cfg.Env)
	s := &Server{
		mux:    http.NewServeMux(),
		config: cfg,
	}

	s.routes()

	return s
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /health", s.handleHealth)
	s.mux.HandleFunc("POST /ingest", s.handleIngest)
}
