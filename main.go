package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/model"
	"github.com/orhanfatih/blog-api/repository"
	"github.com/orhanfatih/blog-api/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load environment variables: %s", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Istanbul", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{}, &model.Post{})

	authStore := repository.NewAuthRepository(db)
	postStore := repository.NewPostRepository(db)
	userStore := repository.NewUserRepository(db)
	srv := server.NewServer(authStore, postStore, userStore)
	g := srv.E.Group("/v1")

	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Blog API")
	})

	srv.RegisterAuthRoutes(g)
	srv.RegisterPostRoutes(g)
	srv.RegisterUserRoutes(g)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("SERVER_PORT is not specified in the .env file")
	}

	if err := srv.E.Start(":" + port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
