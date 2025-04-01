package metrics

import (
	"errors"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server represents an HTTP server for exposing metrics
type Server struct {
	srv     *http.ServeMux
	address string
}

// NewServer creates a new HTTP server for exposing metrics
func NewServer(address string) Server {
	srv := http.NewServeMux()
	srv.Handle("/metrics", promhttp.Handler())

	return Server{
		address: address,
		srv:     srv,
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	srv := http.Server{
		Addr:         s.address,
		Handler:      s.srv,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
		TLSConfig:    nil,
	}

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
