package config

import (
	"log"

	"github.com/joho/godotenv"
)

// The init() function in Go is a special function that allows you to perform initialization tasks before the program starts executing.
func init() {
	// Load environment variables from .env file
	if err := godotenv.Load("local.env"); err != nil {
		log.Printf("Error loading .env file: %s", err)
	}
}
