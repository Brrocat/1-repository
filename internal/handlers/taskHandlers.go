package handlers

import (
	"encoding/json"
	"go.mod/internal/database"
	"go.mod/internal/taskService"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []taskService.Task
	result := database.DB.Find(&tasks)
	if result.Error != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result := database.DB.Create(&task)
	if result.Error != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
