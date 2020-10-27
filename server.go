package main

import (
	_ "go-echo-swagger-sports-api/controllers"
	_ "go-echo-swagger-sports-api/db"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":3000"))
}
