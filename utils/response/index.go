package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error string      `json:"error,omitempty"`
}

func SendResponse(w http.ResponseWriter, statusCode int, data interface{}, meta interface{}, errMessage string) {
	response := APIResponse{
		Data:  data,
		Meta:  meta,
		Error: errMessage,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
