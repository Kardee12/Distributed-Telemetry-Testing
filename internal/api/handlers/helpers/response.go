package helpers

//response.go standarizes how we would have a response

import (
	"encoding/json"
	"net/http"
)

// APIResponse is your standard success shape.
type APIResponse struct {
	Data  any `json:"data,omitempty"`
	Count int `json:"count,omitempty"` // For lists
	Meta  any `json:"meta,omitempty"`  // Optional extras
}

// APIError is your standard error shape.
type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// WriteJSON writes a success response.
func WriteJSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, `{"message":"internal server error","code":500}`, http.StatusInternalServerError)
	}
}

// WriteError writes a standard error response.
func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errResp := APIError{
		Message: message,
		Code:    status,
	}

	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		http.Error(w, `{"message":"internal server error","code":500}`, http.StatusInternalServerError)
	}
}

// Helpers to wrap payloads easily.
func NewAPIResponse(data any) APIResponse {
	return APIResponse{
		Data: data,
	}
}

func NewAPIListResponse(data any, count int) APIResponse {
	return APIResponse{
		Data:  data,
		Count: count,
	}
}
