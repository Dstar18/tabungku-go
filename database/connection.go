package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Load .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Error loading .env file")
	}

	// Connect to DB
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		"tabungdb",
		"5432",
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection established!")
}
