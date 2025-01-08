package routes

import (
	"database/sql"

	"github.com/pump-p/naturemate/handlers"

	"github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	productRoutes := e.Group("/products")

	// // Attach JWT middleware to the product group
	// productRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte("your-secret-key"),
	// }))

	productRoutes.POST("", handlers.CreateProductHandler(db))
	productRoutes.POST("/bulk", handlers.CreateProductsBulkHandler(db))
	productRoutes.GET("/:id", handlers.GetProductByIDHandler(db))
	productRoutes.GET("", handlers.GetAllProductsHandler(db))
	productRoutes.PUT("/:id", handlers.UpdateProductHandler(db))
	productRoutes.DELETE("/:id", handlers.DeleteProductHandler(db))

	//DEVELOPEMENT ONLY
	productRoutes.DELETE("", handlers.DeleteAllProductsHandler(db))
}
