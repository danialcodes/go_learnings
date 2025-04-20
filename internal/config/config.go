package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// all the configuration for the application
type Config struct {
	DatabaseURL string
	ServerPort  string
	Environment string
}

// Load loads configuration from environment variables
func Load() *Config {

	// Load .env file if it exists or exit the application

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, exiting the application")
		os.Exit(1)
	}

	// Check if required environment variables are set
	if os.Getenv("DATABASE_URL") == "" {
		log.Println("DATABASE_URL is not set, exiting the application")
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Default port
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug" // Default mode
	}

	// Set config values
	config := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ServerPort:  port,
		Environment: mode,
	}

	return config
}