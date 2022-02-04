package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

// DummyEndpoint example
// @Description just a dummy endpoint
// @Accept  json
// @Produce  json
// @Success 200 {object} Response	""
// @Router /api/endpoint [get]
func DummyEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	response := Response{
		Message: "Welcome to the API",
	}
	json.NewEncoder(w).Encode(&response)
}
