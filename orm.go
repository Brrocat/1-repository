package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type Task struct {
	gorm.Model
	Message string `json:"message"`
	IsDone  bool   `json:"isDone"`
}

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(w http.ResponseWriter, message string, data interface{}, statusCoda int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCoda)
	json.NewEncoder(w).Encode(ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}
func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ApiResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}
