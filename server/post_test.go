package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/model"
	"github.com/stretchr/testify/assert"
)

func TestHandleCreatePost(t *testing.T) {

	tests := []struct {
		method            string
		route             string
		body              interface{}
		authReq           bool
		cred              *model.LoginRequest
		expectedError     bool
		expectedErrorDesc string
		expectedCode      int
	}{
		{
			// missing auth token
			method:            "POST",
			route:             "/v1/posts/",
			body:              nil,
			authReq:           false,
			cred:              nil,
			expectedError:     true,
			expectedErrorDesc: "You must be logged in to access this resource.",
			expectedCode:      http.StatusUnauthorized,
		},
		{
			// request invalid body field
			method:            "POST",
			route:             "/v1/posts/",
			body:              []byte(`{}`),
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusBadRequest,
		},
		{
			// request invalid body
			method:            "POST",
			route:             "/v1/posts/",
			body:              &model.CreatePostRequest{Title: "Bookstitleissoolong"},
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusBadRequest,
		},
		{
			// success
			method:            "POST",
			route:             "/v1/posts/",
			body:              &model.CreatePostRequest{Title: "Books", Content: "Here are the most influential books of all time ...."},
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusCreated,
		},
		{
			// duplicate title
			method:            "POST",
			route:             "/v1/posts/",
			body:              &model.CreatePostRequest{Title: "Books", Content: "Here are the most influential books of all time ...."},
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		c, resp := makeRequest(test.method, test.route, test.body, test.authReq, test.cred)

		if test.expectedError && test.expectedErrorDesc != "" {
			err := AuthenticateUser(srv.handleCreatePost)(c)
			assert.NotNil(t, err)

			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)

			assert.Equal(t, test.expectedCode, he.Code)
			assert.Equal(t, test.expectedErrorDesc, he.Message)

			continue

		}

		if assert.NoError(t, AuthenticateUser(srv.handleCreatePost)(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}
	}
}

func TestHandleGetPost(t *testing.T) {
	tests := []struct {
		method            string
		route             string
		postId            string
		body              interface{}
		authReq           bool
		cred              *model.LoginRequest
		expectedError     bool
		expectedErrorDesc string
		expectedCode      int
	}{
		{
			// missing auth token
			method:            "GET",
			route:             "/v1/posts/:id",
			postId:            "0",
			body:              nil,
			authReq:           false,
			cred:              nil,
			expectedError:     true,
			expectedErrorDesc: "You must be logged in to access this resource.",
			expectedCode:      http.StatusUnauthorized,
		},
		{
			// invalid postId value
			method:            "GET",
			route:             "/v1/posts/:id",
			postId:            "oops",
			body:              nil,
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusBadRequest,
		},
		{
			// not existing postId
			method:            "GET",
			route:             "/v1/posts/:id",
			postId:            "10000",
			body:              nil,
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusNotFound,
		},
		{
			// success
			method:            "GET",
			route:             "/v1/posts/:id",
			postId:            "1",
			body:              nil,
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     false,
			expectedErrorDesc: "",
			expectedCode:      http.StatusOK,
		},
	}

	for _, test := range tests {
		c, resp := makeRequest(test.method, test.route, test.body, test.authReq, test.cred)
		c.SetPath(test.route)
		c.SetParamNames("id")
		c.SetParamValues(test.postId)
		if test.expectedError && test.expectedErrorDesc != "" {
			err := AuthenticateUser(srv.handleGetPost)(c)
			assert.NotNil(t, err)

			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)

			assert.Equal(t, test.expectedCode, he.Code)
			assert.Equal(t, test.expectedErrorDesc, he.Message)

			continue

		}

		if assert.NoError(t, AuthenticateUser(srv.handleGetPost)(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}

		if !test.expectedError {
			post := model.Post{}
			responseBytes, _ := io.ReadAll(resp.Result().Body)
			_ = json.Unmarshal(responseBytes, &post)
			uintPostId, _ := strconv.ParseUint(test.postId, 0, 64)
			assert.Equal(t, uint(uintPostId), post.ID)
		}
	}
}

func TestHandleUpdatePost(t *testing.T) {

	tests := []struct {
		method            string
		route             string
		postId            string
		body              interface{}
		authReq           bool
		cred              *model.LoginRequest
		expectedError     bool
		expectedErrorDesc string
		expectedCode      int
	}{
		{
			// missing auth token
			method:            "PUT",
			route:             "/v1/posts/:id",
			postId:            "0",
			body:              nil,
			authReq:           false,
			cred:              nil,
			expectedError:     true,
			expectedErrorDesc: "You must be logged in to access this resource.",
			expectedCode:      http.StatusUnauthorized,
		},
		{
			// invalid postId value
			method:            "PUT",
			route:             "/v1/posts/:id",
			postId:            "oops",
			body:              &model.UpdatePostRequest{Title: "Updated Title", Content: "Updated content has no meaning"},
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusBadRequest,
		},
		{
			// not existing postId
			method:            "PUT",
			route:             "/v1/posts/:id",
			postId:            "10000",
			body:              &model.UpdatePostRequest{Title: "Updated Title", Content: "Updated content has no meaning"},
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusNotFound,
		},
		{
			// success
			method:            "PUT",
			route:             "/v1/posts/:id",
			postId:            "1",
			body:              &model.UpdatePostRequest{Title: "Updated Title", Content: "Updated content has no meaning"},
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     false,
			expectedErrorDesc: "",
			expectedCode:      http.StatusOK,
		},
	}

	for _, test := range tests {
		c, resp := makeRequest(test.method, test.route, test.body, test.authReq, test.cred)
		c.SetPath(test.route)
		c.SetParamNames("id")
		c.SetParamValues(test.postId)
		if test.expectedError && test.expectedErrorDesc != "" {
			err := AuthenticateUser(srv.handleUpdatePost)(c)
			assert.NotNil(t, err)

			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)

			assert.Equal(t, test.expectedCode, he.Code)
			assert.Equal(t, test.expectedErrorDesc, he.Message)

			continue

		}

		if assert.NoError(t, AuthenticateUser(srv.handleUpdatePost)(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}

		if !test.expectedError {
			post := model.Post{}
			responseBytes, _ := io.ReadAll(resp.Result().Body)
			_ = json.Unmarshal(responseBytes, &post)
			uintPostId, _ := strconv.ParseUint(test.postId, 0, 64)
			assert.Equal(t, uint(uintPostId), post.ID)
		}
	}

}

func TestHandleDeletePost(t *testing.T) {

	tests := []struct {
		method            string
		route             string
		postId            string
		body              interface{}
		authReq           bool
		cred              *model.LoginRequest
		expectedError     bool
		expectedErrorDesc string
		expectedCode      int
	}{
		{
			// missing auth token
			method:            "DELETE",
			route:             "/v1/posts/:id",
			postId:            "0",
			body:              nil,
			authReq:           false,
			cred:              nil,
			expectedError:     true,
			expectedErrorDesc: "You must be logged in to access this resource.",
			expectedCode:      http.StatusUnauthorized,
		},
		{
			// invalid postId value
			method:            "DELETE",
			route:             "/v1/posts/:id",
			postId:            "oops",
			body:              nil,
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusBadRequest,
		},
		{
			// not existing postId
			method:            "DELETE",
			route:             "/v1/posts/:id",
			postId:            "350",
			body:              nil,
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusBadRequest,
		},
		{
			// success
			method:            "DELETE",
			route:             "/v1/posts/:id",
			postId:            "1",
			body:              nil,
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     false,
			expectedErrorDesc: "",
			expectedCode:      http.StatusNoContent,
		},
	}

	for _, test := range tests {
		c, resp := makeRequest(test.method, test.route, test.body, test.authReq, test.cred)
		c.SetPath(test.route)
		c.SetParamNames("id")
		c.SetParamValues(test.postId)
		if test.expectedError && test.expectedErrorDesc != "" {
			err := AuthenticateUser(srv.handleDeletePost)(c)
			assert.NotNil(t, err)

			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)

			assert.Equal(t, test.expectedCode, he.Code)
			assert.Equal(t, test.expectedErrorDesc, he.Message)

			continue

		}

		if assert.NoError(t, AuthenticateUser(srv.handleDeletePost)(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}
	}
}

func TestHandleExplorePosts(t *testing.T) {

	tests := []struct {
		method            string
		route             string
		authReq           bool
		cred              *model.LoginRequest
		expectedError     bool
		expectedErrorDesc string
		expectedCode      int
	}{
		{
			// missing auth token
			method:            "GET",
			route:             "/v1/posts/",
			authReq:           false,
			cred:              nil,
			expectedError:     true,
			expectedErrorDesc: "You must be logged in to access this resource.",
			expectedCode:      http.StatusUnauthorized,
		},
		{
			// success
			method:            "GET",
			route:             "/v1/posts/",
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     false,
			expectedErrorDesc: "",
			expectedCode:      http.StatusOK,
		},
	}

	for _, test := range tests {
		c, resp := makeRequest(test.method, test.route, nil, test.authReq, test.cred)

		if test.expectedError && test.expectedErrorDesc != "" {
			err := AuthenticateUser(srv.handleExplorePosts)(c)
			assert.NotNil(t, err)

			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)

			assert.Equal(t, test.expectedCode, he.Code)
			assert.Equal(t, test.expectedErrorDesc, he.Message)

			continue

		}

		if assert.NoError(t, AuthenticateUser(srv.handleExplorePosts)(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}

		if !test.expectedError {
			post := []model.Post{}
			responseBytes, _ := io.ReadAll(resp.Result().Body)
			_ = json.Unmarshal(responseBytes, &post)
			log.Println(post)
			// uintPostId, _ := strconv.ParseUint(test.postId, 0, 64)
			// assert.Equal(t, uint(uintPostId), post.ID)
		}
	}
}
