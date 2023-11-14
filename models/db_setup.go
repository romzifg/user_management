package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_connection := os.Getenv("DB_CONNECTION")
	database, err := gorm.Open(mysql.Open(db_connection))
	if err != nil {
		log.Fatal("Error connect to databse")
	}

	database.AutoMigrate(&Role{})
	
	DB = database
}