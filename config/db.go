package config

import (
	"fmt"
	"os"

	"example.com/todoApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected successfully to the database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})
	DB = db

}
