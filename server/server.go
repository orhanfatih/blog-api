package server

import (
	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/repository"
)

type Server struct {
	E *echo.Echo

	store repository.AuthStore
}

func NewServer(store repository.AuthStore) *Server {
	return &Server{E: echo.New(),
		store: store}
}
