package server

import (
	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/repository"
)

type Server struct {
	E *echo.Echo

	authStore repository.AuthStore
	postStore repository.PostStore
	userStore repository.UserStore
}

func NewServer(authStore repository.AuthStore, postStore repository.PostStore, userStore repository.UserStore) *Server {
	return &Server{E: echo.New(),
		authStore: authStore, postStore: postStore, userStore: userStore}
}
