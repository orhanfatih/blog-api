package server

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) RegisterAuthRoutes(g *echo.Group) {
	router := g.Group("/auth")
	router.POST("/register", s.handleRegister)
	router.POST("/login", s.handleLogin)
	router.GET("/logout", s.handleLogout, AuthenticateUser)
}

func (s *Server) handleRegister(c echo.Context) error {
	r := new(models.RegisterRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// pwd match
	if r.Password != r.PasswordConfirm {
		return c.JSON(http.StatusBadRequest, "Passwords do not match!")
	}

	// hash pwd
	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), 10)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// create a User from registerRequest data
	u := models.User{
		Name:      r.Name,
		Email:     r.Email,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}

	if err = s.authStore.CreateUser(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "success")
}

func (s *Server) handleLogin(c echo.Context) error {
	r := new(models.LoginRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var u *models.User
	// query db with email
	u, err := s.authStore.FindUser(u, r.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// check if pwd of loginRequest match with real pwd
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(r.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid login credientials!")
	}

	// create a token
	token, err := CreateToken(u.ID, time.Hour, os.Getenv("JWT_SECRET"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := http.Cookie{
		Name:     "access-token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour).UTC(),
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, "success")
}

func (s *Server) handleLogout(c echo.Context) error {
	cookie := http.Cookie{
		Name:     "access-token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour).UTC(),
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, "logout successful")
}
