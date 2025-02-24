package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mod/internal/database"
	"go.mod/internal/handlers"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Настройка маршрутизатора
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks", handlers.CreateTask).Methods("POST")

	// Запуск сервера
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
