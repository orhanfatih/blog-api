package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/model"
	"github.com/orhanfatih/blog-api/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var srv *Server

func TestMain(m *testing.M) {
	db := mockDatabase()

	db.AutoMigrate(&model.User{}, &model.Post{})

	authStore := repository.NewAuthRepository(db)
	postStore := repository.NewPostRepository(db)
	userStore := repository.NewUserRepository(db)
	srv = NewServer(authStore, postStore, userStore)

	g := srv.E.Group("/v1")

	srv.RegisterAuthRoutes(g)
	srv.RegisterPostRoutes(g)
	srv.RegisterUserRoutes(g)

	exitCode := m.Run()
	teardown(db)

	os.Exit(exitCode)
}

func mockDatabase() *gorm.DB {
	if os.Getenv("CI") == "" {
		err := godotenv.Load("../.env.test")
		if err != nil {
			log.Fatalf("failed to load environment variables: %s", err)
		}
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Istanbul", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		panic(err)
	}

	return db
}

func teardown(db *gorm.DB) {
	migrator := db.Migrator()
	migrator.DropTable(&model.User{}, &model.Post{})
}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool, cred *model.LoginRequest) (echo.Context, *httptest.ResponseRecorder) {
	reqBody, _ := json.Marshal(body)
	req := httptest.NewRequest(method, url, strings.NewReader(string(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	if isAuthenticatedRequest {
		req.AddCookie(bearerToken(cred))
	}

	rec := httptest.NewRecorder()
	c := srv.E.NewContext(req, rec)

	return c, rec
}

func bearerToken(cred *model.LoginRequest) *http.Cookie {

	var u *model.User
	u, _ = srv.authStore.FindUser(u, cred.Email)
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(cred.Password)); err != nil {
		return nil
	}

	token, _ := CreateToken(u.ID, time.Hour, os.Getenv("JWT_SECRET"))

	return &http.Cookie{
		Name:     "access-token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour).UTC(),
		HttpOnly: true,
		Path:     "/",
	}
}
