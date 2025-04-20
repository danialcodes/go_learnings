package api

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danialcodes/crud_api_server/internal/api"
	"github.com/danialcodes/crud_api_server/internal/config"
	"github.com/danialcodes/crud_api_server/internal/db"
)

func Run() {
	// Load configuration
	configs := config.Load()

	// Initialize database
	dbPool, err := db.Initialize(configs)

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	defer dbPool.Close()

	// API server setup
	router := api.SetupRouter(dbPool)
	server := api.NewServer(router, configs.ServerPort)

	// Start the server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", configs.ServerPort)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}