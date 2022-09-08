package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"todoApp/internal/module"
	"todoApp/pkg/router"
)

func main() {
	e := echo.New()

	// Init MongoDB
	module.InitMongoDB()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))

	//Init Router
	router.Init(e)

	// Start Server
	e.Logger.Fatal(e.Start(":9001"))
}
