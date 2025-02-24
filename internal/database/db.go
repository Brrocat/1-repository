package database

import (
	"go.mod/internal/taskService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {

	dsn := "host=localhost user=your_user dbname=your_dbname password=your_password port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to conect to database")
	}
	err = DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}
	log.Println("Database connection and migrated successfully")
}
