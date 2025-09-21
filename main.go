package main

import (
	"snacks-backend/db"
	"snacks-backend/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	// Connect DB
	db.Connect()

	// Setup Echo
	e := echo.New()

	// Routes
	e.GET("/snacks", handlers.GetSnacks)
	e.POST("/snacks", handlers.CreateSnack)
	e.PUT("/snacks/:id", handlers.UpdateSnack)
	e.DELETE("/snacks/:id", handlers.DeleteSnack)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
