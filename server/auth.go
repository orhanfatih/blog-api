package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) RegisterAuthRoutes(g *echo.Group) {
	router := g.Group("/auth")
	router.GET("/signup", s.handleSignup)
}

func (s *Server) handleSignup(c echo.Context) error {
	return c.String(http.StatusOK, "auth-signup")
}
