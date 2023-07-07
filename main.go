package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/api"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load environment variables: %s", err)
	}

	e := echo.New()
	g := e.Group("/api/v1")

	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Blog API")
	})

	api.RegisterAuthRoutes(g)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("SERVER_PORT is not specified in the .env file")
	}

	if err := e.Start(":" + port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
