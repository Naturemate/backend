package main

import (
	"log"

	"github.com/pump-p/naturemate/config"
	"github.com/pump-p/naturemate/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Register routes
	routes.RegisterRoutes(e, db)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
