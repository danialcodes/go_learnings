package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/danialcodes/crud_api_server/internal/models"
	"github.com/danialcodes/crud_api_server/internal/repository"
)

// ProductHandler handles HTTP requests for products
type ProductHandler struct {
	repo *repository.ProductRepository
}

// NewProductHandler creates a new handler instance
func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

// GetProducts handles GET /api/products
func (h *ProductHandler) GetProducts(c *gin.Context) {
	// Parse query parameters for pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	filter := models.ProductFilter{
		Page:  page,
		Limit: limit,
	}

	products, total, err := h.repo.List(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		log.Printf("Database error: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       products,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": (total + limit - 1) / limit,
	})
}

// GetProduct handles GET /api/products/:id
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
			log.Printf("Database error: %v", err)
		}
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct handles POST /api/products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Basic validation
	if product.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name is required"})
		return
	}
	if product.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than zero"})
		return
	}
	if product.Quantity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity cannot be negative"})
		return
	}

	// Create the product
	if err := h.repo.Create(c.Request.Context(), &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		log.Printf("Insert error: %v", err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

// UpdateProduct handles PUT /api/products/:id
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Check if product exists
	exists, err := h.repo.Exists(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check product existence"})
		log.Printf("Database error: %v", err)
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Parse request body
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Basic validation
	if product.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name is required"})
		return
	}
	if product.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than zero"})
		return
	}
	if product.Quantity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity cannot be negative"})
		return
	}

	// Set ID from path parameter
	product.ID = id

	// Update the product
	if err := h.repo.Update(c.Request.Context(), &product); err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			log.Printf("Update error: %v", err)
		}
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct handles DELETE /api/products/:id
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := h.repo.Delete(c.Request.Context(), id); err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			log.Printf("Delete error: %v", err)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}