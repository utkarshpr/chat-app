package models

import (
	"encoding/json"
	"net/http"
)

type GenericResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  bool        `json:status`
}

func ManageResponse(w http.ResponseWriter, String string, code int, data interface{}, status bool) {
	response := GenericResponse{
		Message: String,
		Data:    data,
		Status:  status,
	}
	beautifiedJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(beautifiedJSON)
}
