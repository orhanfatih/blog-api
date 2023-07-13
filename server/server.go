package server

import (
	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/repository"
)

type Server struct {
	E *echo.Echo

	AuthRepository repository.AuthRepository
}

func NewServer() *Server {
	return &Server{E: echo.New()}
}
