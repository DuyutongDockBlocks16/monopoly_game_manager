package handler

import (
	"app-backend/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

// HelloHandler test handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! You've hit %s\n", r.URL.Path)
}

func PropertiesHandler(w http.ResponseWriter, r *http.Request) {
	propertyService := service.NewPropertyService()
	properties, err := propertyService.GetAllProperties()
	if err != nil {
		http.Error(w, "Failed to get properties", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(properties); err != nil {
		http.Error(w, "Failed to encode properties", http.StatusInternalServerError)
	}
}

func (h *Handler) OwnedPropertiesHandler(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("game_id")

	propertyService := service.NewPropertyService()
	properties, err := propertyService.GetOwnedProperties(gameID)

	if err != nil {
		http.Error(w, "Failed to get owned properties", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(properties); err != nil {
		http.Error(w, "Failed to encode properties", http.StatusInternalServerError)
	}
}
