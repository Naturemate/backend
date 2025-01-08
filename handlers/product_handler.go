package handlers

import (
	"database/sql"
	"net/http"

	"github.com/pump-p/naturemate/models"
	"github.com/pump-p/naturemate/utils"

	"github.com/labstack/echo"
)

func CreateProductHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var product models.Product
		if err := c.Bind(&product); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		product.ID = utils.GenerateUUID()

		if err := models.InsertProduct(db, &product); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, product)
	}
}

func CreateProductsBulkHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var products []models.Product
		if err := c.Bind(&products); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		for i := range products {
			products[i].ID = models.GenerateUUID()
		}
		if err := models.InsertProducts(db, products); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, products)
	}
}

func GetProductByIDHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		product, err := models.GetProductByID(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, "Product not found")
			}
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, product)
	}
}

func GetAllProductsHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		products, err := models.GetAllProducts(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, products)
	}
}

func UpdateProductHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		// Parse the incoming JSON into a map for dynamic updates
		var updates map[string]interface{}
		if err := c.Bind(&updates); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// Ensure the ID is not updated
		if _, ok := updates["id"]; ok {
			return c.JSON(http.StatusBadRequest, "Updating the 'id' field is not allowed")
		}

		// Validate fields if necessary (optional, based on your requirements)
		// Example: Check for unsupported fields or validate field values

		// Perform the update
		if err := models.UpdateProduct(db, id, updates); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Product updated successfully",
		})
	}
}

func DeleteProductHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if err := models.DeleteProduct(db, id); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusNoContent)
	}
}

// DEVELOPEMENT ONLY
func DeleteAllProductsHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := models.DeleteAllProducts(db); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusNoContent)
	}
}
