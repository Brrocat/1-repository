package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"net/http"
)

var task string = "Word"

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	result := DB.Find(&tasks)
	if result.Error != nil {
		http.Error(w, "Не удалось получить задачу", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	result := DB.Create(&task)
	if result.Error != nil {
		http.Error(w, "Не удалось создать задачу", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	var updateTask Task

	err := json.NewDecoder(r.Body).Decode(&updateTask)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	var task Task
	result := DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	task.Message = updateTask.Message
	task.IsDone = updateTask.IsDone
	result = DB.Save(&task)
	if result.Error != nil {
		http.Error(w, "Не удалось обновить задачу", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	var task Task
	result := DB.Delete(&task, id)
	if result.Error != nil {
		http.Error(w, "Не удалось удалить задачу", http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
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
