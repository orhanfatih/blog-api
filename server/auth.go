package server

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) RegisterAuthRoutes(g *echo.Group) {
	router := g.Group("/auth")
	router.POST("/register", s.handleRegister)
	router.POST("/login", s.handleLogin)
}

func (s *Server) handleRegister(c echo.Context) error {
	r := new(models.RegisterRequest)
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// pwd match
	if r.Password != r.PasswordConfirm {
		return c.String(http.StatusBadRequest, "Passwords are not match!")
	}

	// hash pwd
	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), 10)
	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())
	}

	// create User from registerRequest data
	u := models.User{
		Name:      r.Name,
		Email:     r.Email,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}

	if err = s.store.CreateUser(&u); err != nil {
		return c.String(http.StatusBadGateway, "something went wrong")
	}

	return c.String(http.StatusOK, "success")
}

func (s *Server) handleLogin(c echo.Context) error {
	r := new(models.LoginRequest)
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	var u *models.User
	// query db with email
	u, err := s.store.FindUser(u, r.Email)
	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())

	}
	// check if pwd of loginRequest match with real pwd
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(r.Password)); err != nil {
		return c.String(http.StatusBadRequest, "wrong password!")
	}

	// now create a token, return to user
	return c.String(http.StatusOK, "successful login")

	// return nil
}
