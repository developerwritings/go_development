// handler/product_handler.go
package handler

import (
	"encoding/json"
	"net/http"

)

// ProductHandler handles product-related HTTP requests
type ProductHandler struct{}

// GetProducts handles GET request to fetch all products
func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// Your logic to fetch all products from the database
	dbHandler := &db.SQLDBHandler{} // Instantiate your DB handler
	products, err := dbHandler.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, products)
}

// GetProduct handles GET request to fetch a single product by ID
func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Your logic to fetch a single product by ID from the database
	dbHandler := &db.SQLDBHandler{} // Instantiate your DB handler

	productID := r.URL.Query().Get("id")
	product, err := dbHandler.GetProductByID(productID)
	if err != nil {
		http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, product)
}

// CreateProduct handles POST request to create a new product
func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Your logic to create a new product in the database
	dbHandler := &db.SQLDBHandler{} // Instantiate your DB handler

	var productData interface{}
	err := json.NewDecoder(r.Body).Decode(&productData)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	err = dbHandler.CreateProduct(productData)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	respondWithMessage(w, http.StatusCreated, "Product created successfully")
}

// UpdateProduct handles PUT request to update an existing product
func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Your logic to update an existing product in the database
	dbHandler := &db.SQLDBHandler{} // Instantiate your DB handler

	productID := r.URL.Query().Get("id")

	var productData interface{}
	err := json.NewDecoder(r.Body).Decode(&productData)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	err = dbHandler.UpdateProduct(productID, productData)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
	respondWithMessage(w, http.StatusOK, "Product updated successfully")
}

// DeleteProduct handles DELETE request to delete an existing product
func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Your logic to delete an existing product from the database
	dbHandler := &db.SQLDBHandler{} // Instantiate your DB handler

	productID := r.URL.Query().Get("id")

	err := dbHandler.DeleteProduct(productID)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}
	respondWithMessage(w, http.StatusOK, "Product deleted successfully")
}

// Helper function to respond with JSON data
func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode response data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// Helper function to respond with a plain text message
func respondWithMessage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}
