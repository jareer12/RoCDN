package main

import (
	"fmt"
	"rblx/database"
	"rblx/routes"
	"time"

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

	go func() {
		for true {
			data := hs.Data

			for i := 0; i < len(data); i++ {
				r := data[i]

				if time.Now().UnixMilli() >= r.Timestamp {
					database.Remove(&hs, r.TargetId)
				}
			}

			time.Sleep(time.Second * 5)
			fmt.Println(fmt.Sprintf(`Goroutine Checked %v Cached Images`, len(data)))
		}
	}()

	// Start server
	e.Logger.Fatal(e.Start("127.0.0.1:1313"))
}
