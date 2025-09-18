package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ZiplEix/PDF-tools/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Bienvenue sur l'API")
	})

	routes.SetupRoutes(e)

	fmt.Println("Server is running on https://0.0.0.0:8080")
	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
