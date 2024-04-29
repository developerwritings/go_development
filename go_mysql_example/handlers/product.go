// handler/product_handler.go
package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ProductHandler handles product-related HTTP requests
type ProductHandler struct {
	DB *sql.DB
}

// NewProductHandler initializes a new ProductHandler
func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{DB: db}
}

// GetProducts handles GET request to fetch all products
func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// Your logic to fetch all products from the database
}

// GetProduct handles GET request to fetch a single product by ID
func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Your logic to fetch a single product by ID from the database
}

// CreateProduct handles POST request to create a new product
func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Your logic to create a new product in the database
}

// UpdateProduct handles PUT request to update an existing product
func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Your logic to update an existing product in the database
}

// DeleteProduct handles DELETE request to delete an existing product
func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Your logic to delete an existing product from the database
}
