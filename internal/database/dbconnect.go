package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mysterybee07/jwt-auth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file. Err: %s", err)
		return
	}
}

func DBconn() {
	LoadEnvironment()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error Connecting to database %s", err)
		return
	}
	log.Println("Connected to database successfully........")

	DB = db

	db.AutoMigrate(&models.User{}) // we are going to create a models.go file for the User Model.
}
