package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/model"
)

func (s *Server) RegisterUserRoutes(g *echo.Group) {
	router := g.Group("/user")
	router.Use(AuthenticateUser)
	router.GET("/me", s.handleGetMe)
	router.PATCH("/", s.handleUpdateProfile)
	router.DELETE("/", s.handleDeleteProfile)
}

func (s *Server) handleGetMe(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return RespondWithError(c, http.StatusInternalServerError, "User ID not found in context")
	}

	user, err := s.userStore.FindUser(userID)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	u := model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return RespondWithJSON(c, http.StatusOK, u)
}

func (s *Server) handleUpdateProfile(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return RespondWithError(c, http.StatusInternalServerError, "User ID not found in context")
	}

	r := new(model.User)
	if err := c.Bind(r); err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	user, err := s.userStore.UpdateUser(userID, r)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	u := model.UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return RespondWithJSON(c, http.StatusOK, u)
}

func (s *Server) handleDeleteProfile(c echo.Context) error {
	e := new(model.User)
	if err := c.Bind(e); err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}
	userID, ok := c.Get("userID").(int)
	if !ok {
		return RespondWithError(c, http.StatusInternalServerError, "User ID not found in context")
	}
	e.ID = uint(userID)

	if err := s.userStore.DeleteUser(e); err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/v1/auth/logout")
}
