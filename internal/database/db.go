package database

import (
	"go.mod/internal/taskService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() (*gorm.DB, error) { // Теперь функция возвращает *gorm.DB и ошибку
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
		return nil, err
	}

	err = db.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
		return nil, err
	}

	log.Println("Database connection and migration successful")
	return db, nil
}
