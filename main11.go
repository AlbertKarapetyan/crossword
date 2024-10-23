// main.go
package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	// initDI "crossword/internal/app"
	// "crossword/internal/database/models"
	// "fmt"
)

func main() {
	godotenv.Load()
	// Check for the environment variable
	dbHost := os.Getenv("POSTGRESQL_DB_HOST")

	if dbHost == "" {
		fmt.Println("Environment variable POSTGRESQL_DB_HOST is not set or is empty")
	} else {
		fmt.Printf("POSTGRESQL_DB_HOST: %s\n", dbHost)
	}

	// userService, err := initDI.InitializeService()
	// if err != nil {
	// 	fmt.Println("Failed to initialize the service:", err)
	// 	return
	// }

	// // Example usage of userService
	// err = userService.CreateUser(&models.User{ID: 1, ExternalID: 111})
	// if err != nil {
	// 	fmt.Println("Error creating user:", err)
	// } else {
	// 	fmt.Println("User created successfully!")
	// }

	// err = userService.DeleteUser(1)
	// if err != nil {
	// 	fmt.Println("Error creating user:", err)
	// } else {
	// 	fmt.Println("User created successfully!")
	// }
}
