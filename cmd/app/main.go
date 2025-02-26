package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mod/internal/database"
	"go.mod/internal/handlers"
	"go.mod/internal/taskService"
	"log"
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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/tasks", taskHandler.GetTask)
	e.POST("/api/tasks", taskHandler.PostTask)
	e.PATCH("/api/tasks/:id", taskHandler.PatchTask)
	e.DELETE("/api/tasks/:id", taskHandler.DeleteTask)

	e.Logger.Fatal(e.Start(":8083"))

}
