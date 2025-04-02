package metrics

import (
	"errors"
	"net/http"
	"time"

	"github.com/amirhnajafiz/nfs-metrics-exporter/pkg/hashing"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server represents an HTTP server for exposing metrics
type Server struct {
	srv     *http.ServeMux
	address string
}

// NewServer creates a new HTTP server for exposing metrics
func NewServer(address string, secret string) Server {
	srv := http.NewServeMux()
	srv.Handle("/metrics", promhttp.Handler())
	srv.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	srv.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	srv.HandleFunc("/valz", func(w http.ResponseWriter, r *http.Request) {
		md5 := hashing.MD5([]byte(secret))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(md5))
	})

	return Server{
		address: ":" + address,
		srv:     srv,
	}
}

// Start starts the HTTP server
func (s Server) Start() error {
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
