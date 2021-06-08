package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"brsrker.com/diamond/proxyserver/internal/logger"
	v1 "brsrker.com/diamond/proxyserver/internal/server/v1"
)

// Server is a base server configuration.
type Server struct {
	server *http.Server
}

const TAG = "server"

// New inicialize a new server with configuration.
func New(port string) (*Server, error) {
	r := chi.NewRouter()

	// API routes version 1.
	r.Mount("/api/v1", v1.New())

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

// Close server resources.
func (serv *Server) Close() error {
	// TODO: add resource closure.
	return nil
}

// Start the server.
func (serv *Server) Start() {
	logger.Info(TAG, fmt.Sprintf("Server running on http://localhost%s", serv.server.Addr))
	log.Fatal(serv.server.ListenAndServe())
}
