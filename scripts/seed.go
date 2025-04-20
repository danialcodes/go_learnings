package main

import (
	"context"
	"log"

	"github.com/danialcodes/crud_api_server/internal/config"
	"github.com/danialcodes/crud_api_server/internal/db"
	"github.com/danialcodes/crud_api_server/internal/models"
	"github.com/danialcodes/crud_api_server/internal/repository"
)

// Seed data
var seedProducts = []models.Product{
	{Name: "Laptop", Price: 1299.99, Quantity: 10},
	{Name: "Smartphone", Price: 699.99, Quantity: 20},
	{Name: "Headphones", Price: 199.99, Quantity: 30},
	{Name: "Monitor", Price: 349.99, Quantity: 5},
	{Name: "Keyboard", Price: 89.99, Quantity: 15},
}

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	dbPool, err := db.Initialize(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbPool.Close()

	// Create repository
	productRepo := repository.NewProductRepository(dbPool)

	// Check if seeding is required
	ctx := context.Background()
	_, count, err := productRepo.List(ctx, models.ProductFilter{Page: 1, Limit: 1})
	if err != nil {
		log.Fatalf("Failed to check database: %v", err)
	}

	if count > 0 {
		log.Println("Database already contains products, skipping seed")
		log.Printf("Current product count: %d", count)
		return
	}

	// Seed data
	log.Println("Seeding database...")

	for _, product := range seedProducts {
		if err := productRepo.Create(ctx, &product); err != nil {
			log.Fatalf("Failed to seed product: %v", err)
		}
	}

	log.Println("Database seeding completed successfully")
}