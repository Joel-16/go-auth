package config

import (
	"encoding/json"
	"net/http"
)

// MessageResponse is a struct for wrapping string responses
type MessageResponse struct {
	Message string `json:"message"`
}

// Respond sends JSON responses, handling both strings and structs
func Respond(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Check if data is a string, wrap it in a struct
	if message, ok := data.(string); ok {
		data = MessageResponse{Message: message}
	}

	// Encode and send the response
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
