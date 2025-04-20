package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/danialcodes/crud_api_server/internal/models"
)

// ProductRepository handles database operations for products
type ProductRepository struct {
	db *pgxpool.Pool
}

// NewProductRepository creates a new repository instance
func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

// List retrieves products with pagination
func (r *ProductRepository) List(ctx context.Context, filter models.ProductFilter) ([]models.Product, int, error) {

	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	offset := (filter.Page - 1) * filter.Limit

	// Get products with pagination
	rows, err := r.db.Query(ctx, `
		SELECT *
		FROM products
		ORDER BY id
		LIMIT $1 OFFSET $2
	`, filter.Limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Parse the rows into products
	var products []models.Product //products slice to hold the results

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	// Get total count
	var totalCount int
	err = r.db.QueryRow(ctx, "SELECT COUNT(*) FROM products").Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}

// GetByID retrieves a product by ID
func (r *ProductRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	var p models.Product
	
	err := r.db.QueryRow(ctx, `
		SELECT id, name, price, quantity, created_at, updated_at
		FROM products
		WHERE id = $1
	`, id).Scan(&p.ID, &p.Name, &p.Price, &p.Quantity, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &p, nil
}

// Create inserts a new product
func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO products (name, price, quantity)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`, product.Name, product.Price, product.Quantity).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
}

// Update modifies an existing product
func (r *ProductRepository) Update(ctx context.Context, product *models.Product) error {
	err := r.db.QueryRow(ctx, `
		UPDATE products
		SET name = $1, price = $2, quantity = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING created_at, updated_at
	`, product.Name, product.Price, product.Quantity, product.ID).Scan(&product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("product not found")
		}
		return err
	}

	return nil
}

// Delete removes a product by ID
func (r *ProductRepository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(ctx, "DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("product not found")
	}

	return nil
}

// Exists checks if a product exists by ID
func (r *ProductRepository) Exists(ctx context.Context, id int) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM products WHERE id = $1)", id).Scan(&exists)
	return exists, err
}