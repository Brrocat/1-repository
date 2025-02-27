package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mod/internal/database"
	"go.mod/internal/handlers"
	"go.mod/internal/taskService"
	"go.mod/internal/userService"
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

	usersRepo := userService.NewRepository(db)
	usersService := userService.NewService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Роуты для задач
	e.GET("/api/tasks", taskHandler.GetTask)
	e.POST("/api/tasks", taskHandler.PostTask)
	e.PATCH("/api/tasks/:id", taskHandler.PatchTask)
	e.DELETE("/api/tasks/:id", taskHandler.DeleteTask)

	//Роуты для пользователей
	e.GET("/api/users", usersHandler.GetUsers)
	e.POST("/api/users", usersHandler.PostUser)
	e.PATCH("api/users/:id", usersHandler.PatchUserId)
	e.DELETE("/api/users/:id", usersHandler.DeleteUser)

	e.Logger.Fatal(e.Start(":8083"))

}
