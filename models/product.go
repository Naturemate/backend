package models

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
)

type Product struct {
	ID              string   `json:"id" db:"id"`
	Name            string   `json:"name" db:"name"`
	Brand           string   `json:"brand" db:"brand"`
	Category        string   `json:"category" db:"category"`
	Price           float64  `json:"price" db:"price"`
	KeyBenefit      string   `json:"key_benefit" db:"key_benefit"`
	CapsuleQuantity *int     `json:"capsule_quantity" db:"capsule_quantity"`
	CapsuleType     string   `json:"capsule_type" db:"capsule_type"`
	Weight          *float64 `json:"weight" db:"weight"`
	SupplementFact  []string `json:"supplement_fact" db:"supplement_fact"`
	Dosage          string   `json:"dosage" db:"dosage"`
	ImageURL        *string  `json:"image_url" db:"image_url"`
	FDA             *string  `json:"fda" db:"fda"`
	FDAURL          *string  `json:"fda_url" db:"fda_url"`
}

// GenerateUUID generates a new UUID string
func GenerateUUID() string {
	return uuid.New().String()
}

// InsertProduct inserts a single product into the database
func InsertProduct(db *sql.DB, product *Product) error {
	query := `INSERT INTO products (id, name, brand, category, price, key_benefit, capsule_quantity, capsule_type, weight, supplement_fact, dosage, image_url, fda, fda_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	supplementFactJSON, _ := json.Marshal(product.SupplementFact)
	_, err := db.Exec(query, product.ID, product.Name, product.Brand, product.Category, product.Price, product.KeyBenefit, product.CapsuleQuantity, product.CapsuleType, product.Weight, supplementFactJSON, product.Dosage, product.ImageURL, product.FDA, product.FDAURL)
	return err
}

// InsertProducts inserts multiple products into the database
func InsertProducts(db *sql.DB, products []Product) error {
	for _, product := range products {
		if err := InsertProduct(db, &product); err != nil {
			return err
		}
	}
	return nil
}

// GetProductByID retrieves a product by its ID from the database
func GetProductByID(db *sql.DB, id string) (*Product, error) {
	query := `SELECT id, name, brand, category, price, key_benefit, capsule_quantity, capsule_type, weight, supplement_fact, dosage, image_url, fda, fda_url FROM products WHERE id = ?`
	var product Product
	var supplementFactJSON []byte
	row := db.QueryRow(query, id)
	if err := row.Scan(&product.ID, &product.Name, &product.Brand, &product.Category, &product.Price, &product.KeyBenefit, &product.CapsuleQuantity, &product.CapsuleType, &product.Weight, &supplementFactJSON, &product.Dosage, &product.ImageURL, &product.FDA, &product.FDAURL); err != nil {
		return nil, err
	}
	json.Unmarshal(supplementFactJSON, &product.SupplementFact)
	return &product, nil
}

// GetAllProducts retrieves all products from the database
func GetAllProducts(db *sql.DB) ([]Product, error) {
	query := `SELECT id, name, brand, category, price, key_benefit, capsule_quantity, capsule_type, weight, supplement_fact, dosage, image_url, fda, fda_url FROM products`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		var supplementFactJSON []byte
		if err := rows.Scan(&product.ID, &product.Name, &product.Brand, &product.Category, &product.Price, &product.KeyBenefit, &product.CapsuleQuantity, &product.CapsuleType, &product.Weight, &supplementFactJSON, &product.Dosage, &product.ImageURL, &product.FDA, &product.FDAURL); err != nil {
			return nil, err
		}
		json.Unmarshal(supplementFactJSON, &product.SupplementFact)
		products = append(products, product)
	}
	return products, nil
}

// UpdateProduct updates an existing product in the database
func UpdateProduct(db *sql.DB, id string, updates map[string]interface{}) error {
	// Check if the product exists
	existingProduct, err := GetProductByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		}
		return err
	}

	// Merge updates into the existing product
	updatedProduct := *existingProduct
	for key, value := range updates {
		switch key {
		case "name":
			updatedProduct.Name = value.(string)
		case "brand":
			updatedProduct.Brand = value.(string)
		case "category":
			updatedProduct.Category = value.(string)
		case "price":
			updatedProduct.Price = value.(float64)
		case "key_benefit":
			updatedProduct.KeyBenefit = value.(string)
		case "capsule_quantity":
			if v, ok := value.(int); ok {
				updatedProduct.CapsuleQuantity = &v
			} else {
				updatedProduct.CapsuleQuantity = nil
			}
		case "capsule_type":
			updatedProduct.CapsuleType = value.(string)
		case "weight":
			if v, ok := value.(float64); ok {
				updatedProduct.Weight = &v
			} else {
				updatedProduct.Weight = nil
			}
		case "supplement_fact":
			updatedProduct.SupplementFact = value.([]string)
		case "dosage":
			updatedProduct.Dosage = value.(string)
		case "image_url":
			if v, ok := value.(string); ok {
				updatedProduct.ImageURL = &v
			} else {
				updatedProduct.ImageURL = nil
			}
		case "fda":
			if v, ok := value.(string); ok {
				updatedProduct.FDA = &v
			} else {
				updatedProduct.FDA = nil
			}
		case "fda_url":
			if v, ok := value.(string); ok {
				updatedProduct.FDAURL = &v
			} else {
				updatedProduct.FDAURL = nil
			}
		}
	}

	// Update the product in the database
	query := `UPDATE products SET name = ?, brand = ?, category = ?, price = ?, key_benefit = ?, capsule_quantity = ?, capsule_type = ?, weight = ?, supplement_fact = ?, dosage = ?, image_url = ?, fda = ?, fda_url = ? WHERE id = ?`
	supplementFactJSON, _ := json.Marshal(updatedProduct.SupplementFact)
	_, err = db.Exec(query, updatedProduct.Name, updatedProduct.Brand, updatedProduct.Category, updatedProduct.Price, updatedProduct.KeyBenefit, updatedProduct.CapsuleQuantity, updatedProduct.CapsuleType, updatedProduct.Weight, supplementFactJSON, updatedProduct.Dosage, updatedProduct.ImageURL, updatedProduct.FDA, updatedProduct.FDAURL, id)
	return err
}

// DeleteProduct deletes a product by its ID from the database
func DeleteProduct(db *sql.DB, id string) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

// DeleteAllProducts deletes all products from the database (Development only)
func DeleteAllProducts(db *sql.DB) error {
	query := `DELETE FROM products`
	_, err := db.Exec(query)
	return err
}
