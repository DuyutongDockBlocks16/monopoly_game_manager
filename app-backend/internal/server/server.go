package server

import (
	"app-backend/internal/handler"
	"net/http"
)

// Start HTTP server
func Start(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HelloHandler)
	mux.HandleFunc("/properties", handler.PropertiesHandler)
	mux.HandleFunc("/properties/owned", handler.OwnedPropertiesHandler)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}
