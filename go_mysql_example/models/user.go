// models/product.go
package models

// Product represents a product in the e-commerce application
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	// Add more fields as needed
}
