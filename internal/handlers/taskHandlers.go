package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mod/internal/taskService"
	"net/http"
)

type TaskHandler struct {
	service *taskService.Service
}

func NewTaskHandler(service *taskService.Service) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err = h.service.CreateTask(&task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updateTask taskService.Task

	err := json.NewDecoder(r.Body).Decode(&updateTask)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	task, err := h.service.UpdateTask(id, &updateTask)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.service.DeleteTask(id)
	if err != nil {
		http.Error(w, "Failed to Delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
