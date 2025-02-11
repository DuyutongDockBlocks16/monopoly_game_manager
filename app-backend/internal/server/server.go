package server

import (
	"net/http"
	"app-backend/internal/handler"
)

// Start HTTP server
func Start(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HelloHandler)
	mux.HandleFunc("/properties", handler.PropertiesHandler)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}
