package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/models"
)

func (s *Server) RegisterPostRoutes(g *echo.Group) {
	router := g.Group("/posts")
	router.Use(AuthenticateUser)
	router.POST("", s.handleCreatePost)
	router.GET("/:id", s.handleGetPost)
	router.PUT("/:id", s.handleUpdatePost)
	router.DELETE("/:id", s.handleDeletePost)
}

func (s *Server) handleCreatePost(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "User ID not found in context")
	}

	// bind given post details to createPostRequest model
	r := new(models.CreatePostRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// create a Post
	p := models.Post{
		UserID:    uint(userID),
		Title:     r.Title,
		Content:   r.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.postStore.CreatePost(&p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, p)
}

func (s *Server) handleGetPost(c echo.Context) error {

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid postid format, should be int")
	}

	var post models.Post
	p, err := s.postStore.FindPost(&post, postID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}
func (s *Server) handleUpdatePost(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "User ID not found in context")
	}
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid postid format, should be int")
	}

	r := new(models.UpdatePostRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var post *models.Post
	post, err = s.postStore.FindPost(post, postID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	p := models.Post{
		UserID:    uint(userID),
		Title:     r.Title,
		Content:   r.Content,
		UpdatedAt: time.Now(),
	}

	if err = s.postStore.UpdatePost(post, &p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}
func (s *Server) handleDeletePost(c echo.Context) error {

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid postid format, should be int")
	}

	if err = s.postStore.DeletePost(postID); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
