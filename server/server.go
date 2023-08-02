package server

import (
	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/repository"
)

type Server struct {
	E *echo.Echo

	authStore repository.AuthStore
	postStore repository.PostStore
}

func NewServer(authStore repository.AuthStore, postStore repository.PostStore) *Server {
	return &Server{E: echo.New(),
		authStore: authStore, postStore: postStore}
}
