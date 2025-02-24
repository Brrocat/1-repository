package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.mod/internal/database"
	"go.mod/internal/handlers"
	"go.mod/internal/taskService"
	"log"
	"net/http"
)

func main() {
	// Инициализация базы данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Настройка маршрутизатора
	repo := taskService.NewRepository(db)
	service := taskService.NewService(repo)
	taskHandler := handlers.NewTaskHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", taskHandler.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", taskHandler.PatchTask).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	// Запуск сервера
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
