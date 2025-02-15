package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	result := DB.Find(&tasks)
	if result.Error != nil {
		ErrorResponse(w, "Failed to fetch task", http.StatusInternalServerError)
		return
	}
	SuccessResponse(w, "Task fetched successfully", tasks, http.StatusOK)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if task.Message == "" {
		ErrorResponse(w, "Message cannot by empty", http.StatusBadRequest)
		return
	}
	result := DB.Create(&task)
	if result.Error != nil {
		ErrorResponse(w, "Failed to create task", http.StatusCreated)
		return
	}
	SuccessResponse(w, "Task created successfully", task, http.StatusCreated)
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	var updateTask Task

	err := json.NewDecoder(r.Body).Decode(&updateTask)
	if err != nil {
		ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if updateTask.Message == "" {
		ErrorResponse(w, "Message cannot by empty", http.StatusBadRequest)
		return
	}

	var task Task
	result := DB.First(&task, id)
	if result.Error != nil {
		ErrorResponse(w, "Task not found", http.StatusNotFound)
		return
	}
	task.Message = updateTask.Message
	task.IsDone = updateTask.IsDone

	result = DB.Save(&task)
	if result.Error != nil {
		ErrorResponse(w, "Failed to update task", http.StatusInternalServerError)
		return
	}
	SuccessResponse(w, "Task updated successfully", task, http.StatusOK)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	var task Task
	result := DB.Delete(&task, id)
	if result.Error != nil {
		ErrorResponse(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		ErrorResponse(w, "Task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {

	InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetHandler).Methods("GET")
	router.HandleFunc("/api/tasks", PostHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", DeleteHandler).Methods("DELETE")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe("localhost:8080", router)
}
