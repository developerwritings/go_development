// db/database.go
package db

import (
	"errors"

	"gorm.io/gorm"
)

// Product represents the product model
type Product struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

// GormDBHandler implements DBHandler interface using GORM
type GormDBHandler struct {
	DB *gorm.DB
}

// GetAllProducts retrieves all products from the database
func (dbh *GormDBHandler) GetAllProducts() ([]string, error) {
	var products []Product
	result := dbh.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	var productNames []string
	for _, p := range products {
		productNames = append(productNames, p.Name)
	}

	return productNames, nil
}

// GetProductByID retrieves a product by ID from the database
func (dbh *GormDBHandler) GetProductByID(id uint) (string, error) {
	var product Product
	result := dbh.DB.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", errors.New("product not found")
		}
		return "", result.Error
	}

	return product.Name, nil
}

// CreateProduct creates a new product in the database
func (dbh *GormDBHandler) CreateProduct(name string) error {
	product := Product{Name: name}
	result := dbh.DB.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateProduct updates an existing product in the database
func (dbh *GormDBHandler) UpdateProduct(id uint, name string) error {
	var product Product
	result := dbh.DB.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return result.Error
	}

	product.Name = name
	result = dbh.DB.Save(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteProduct deletes a product from the database
func (dbh *GormDBHandler) DeleteProduct(id uint) error {
	result := dbh.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
