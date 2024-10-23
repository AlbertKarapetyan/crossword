package main

import (
	"crossword/internal/database"
	model "crossword/internal/database/models"
	"log"
	"os"

	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load the appropriate .env file based on the environment
	env := os.Getenv("APP_ENV")

	if env == "development" {
		err = godotenv.Load("../../.env.dev") // Load .env.dev in development
		fmt.Println("Loaded .env.dev for development")
	} else {
		err = godotenv.Load("../../.env") // Load .env for production
		fmt.Println("Loaded .env for production")
	}

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := database.ProvideDB()
	if err != nil {
		fmt.Println("Error on update database", err)
	}

	// Automigrate schema
	err = db.AutoMigrate(&model.Category{})
	if err != nil {
		fmt.Println("Error on Category", err)
	}
	err = db.AutoMigrate(&model.Word{})
	if err != nil {
		fmt.Println("Error on Word", err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("Error on User", err)
	}

}
