package controllers

import (
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
		status := map[string]string{"status": "on"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
}