package config

import (
	"fmt"
	"log"
	"os"

	"auth/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var JWT_SECRET string
var PORT string

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	PORT = os.Getenv("PORT")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	fmt.Println("Database connected!")

	// Auto-migrate models
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Blog{})
	DB.AutoMigrate(&models.Comment{})
}
