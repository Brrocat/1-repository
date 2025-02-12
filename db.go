package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключится к базе данных", err)
	}
	err = DB.AutoMigrate(&Task{})
	if err != nil {
		log.Fatal("Не удалось перенести базу данных", err)
	}
	log.Println("База данных подключена и успешно перенесена")
}
