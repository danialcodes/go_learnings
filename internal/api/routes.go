package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/danialcodes/crud_api_server/internal/api/handlers"
	"github.com/danialcodes/crud_api_server/internal/api/middleware"
	"github.com/danialcodes/crud_api_server/internal/repository"
)

// SetupRouter configures the Gin router with routes and middleware
func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.Cors())
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())

	// Create repositories
	productRepo := repository.NewProductRepository(db)

	// Create handlers
	productHandler := handlers.NewProductHandler(productRepo)

	// API routes
	api := router.Group("/api/v1")

	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
			products.GET("/:id", productHandler.GetProduct)
			products.POST("", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}
	}


	return router
}

// NewServer creates a new HTTP server with the given router
func NewServer(router *gin.Engine, port string) *http.Server {
	return &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}