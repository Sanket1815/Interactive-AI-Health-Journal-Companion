package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type APIError struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Errors  []ValidationError  `json:"errors,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteSuccess(w http.ResponseWriter, message string, data interface{}) {
	WriteJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func WriteCreated(w http.ResponseWriter, message string, data interface{}) {
	WriteJSON(w, http.StatusCreated, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, APIResponse{
		Success: false,
		Error:   message,
	})
}

func WriteValidationError(w http.ResponseWriter, errors []ValidationError) {
	WriteJSON(w, http.StatusBadRequest, APIError{
		Code:    http.StatusBadRequest,
		Message: "Validation failed",
		Errors:  errors,
	})
}