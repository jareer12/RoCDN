package main

import (
	"rblx/database"
	"rblx/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Headshot store
	hs := database.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", routes.PrimaryRoute)
	e.GET("/headshot/:userId", routes.Headshot(&hs))

	// Start server
	e.Logger.Fatal(e.Start("127.0.0.1:1313"))
}
