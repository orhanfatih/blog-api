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
	router.POST("/", s.handleCreatePost)
	router.GET("/:id", s.handleGetPost)
	router.PUT("/:id", s.handleUpdatePost)
	router.DELETE("/:id", s.handleDeletePost)
	router.GET("/", s.handleExplorePosts)
}

func (s *Server) handleCreatePost(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return RespondWithError(c, http.StatusInternalServerError, "User ID not found in context")
	}

	// bind given post details to createPostRequest model
	r := new(models.CreatePostRequest)
	if err := c.Bind(r); err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
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
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	return RespondWithJSON(c, http.StatusCreated, p)
}

func (s *Server) handleGetPost(c echo.Context) error {

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, "Provide postid")
	}

	var post models.Post
	p, err := s.postStore.FindPost(&post, postID)
	if err != nil {
		return RespondWithError(c, http.StatusNotFound, err.Error())
	}

	return RespondWithJSON(c, http.StatusOK, p)
}

func (s *Server) handleUpdatePost(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return RespondWithError(c, http.StatusInternalServerError, "User ID not found in context")
	}
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, "Provide postid")
	}

	r := new(models.UpdatePostRequest)
	if err := c.Bind(r); err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	var post *models.Post
	post, err = s.postStore.FindPost(post, postID)
	if err != nil {
		return RespondWithError(c, http.StatusNotFound, err.Error())
	}

	p := models.Post{
		UserID:    uint(userID),
		Title:     r.Title,
		Content:   r.Content,
		UpdatedAt: time.Now(),
	}

	if err = s.postStore.UpdatePost(post, &p); err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	return RespondWithJSON(c, http.StatusOK, p)
}

func (s *Server) handleDeletePost(c echo.Context) error {

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, "Provide postid")
	}

	if err = s.postStore.DeletePost(postID); err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	return RespondWithJSON(c, http.StatusNoContent, nil)
}

func (s *Server) handleExplorePosts(c echo.Context) error {
	pageStr := c.QueryParams().Get("page")
	limitStr := c.QueryParams().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit

	posts, err := s.postStore.FindPosts(limit, offset)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	return RespondWithJSON(c, http.StatusOK, posts)
}
