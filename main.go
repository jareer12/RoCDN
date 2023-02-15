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
	// Avatar
	av := database.New()

	// Headshot store
	hs := database.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", routes.PrimaryRoute)
	e.GET("/avatar/:userId", routes.Avatar(&av))
	e.GET("/headshot/:userId", routes.Headshot(&hs))

	// Check & Remove Old Cached Images and Free Memory
	go func() {
		for true {
			for i := 0; i < len(hs.Data); i++ {
				r := hs.Data[i]

				if time.Now().UnixMilli() >= r.Timestamp {
					database.Remove(&hs, r.TargetId)
				}
			}

			for i := 0; i < len(av.Data); i++ {
				r := av.Data[i]

				if time.Now().UnixMilli() >= r.Timestamp {
					database.Remove(&av, r.TargetId)
				}
			}

			fmt.Println(av, hs)
			time.Sleep(time.Second * 5)
		}
	}()

	// Start server
	e.Logger.Fatal(e.Start("127.0.0.1:1313"))
}
