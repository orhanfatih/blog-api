package server

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleGetMe(t *testing.T) {

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
			method:            "GET",
			route:             "/v1/user/me",
			body:              nil,
			authReq:           false,
			cred:              nil,
			expectedError:     true,
			expectedErrorDesc: "You must be logged in to access this resource.",
			expectedCode:      http.StatusUnauthorized,
		},
		{
			// success
			method:            "GET",
			route:             "/v1/user/me",
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

		if test.expectedError {
			err := AuthenticateUser(srv.handleGetMe)(c)
			assert.NotNil(t, err)

			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)

			assert.Equal(t, test.expectedCode, he.Code)
			assert.Equal(t, test.expectedErrorDesc, he.Message)

			continue

		}

		if assert.NoError(t, AuthenticateUser(srv.handleGetMe)(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}
	}

}

func TestHandleUpdateProfile(t *testing.T) {

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
			method:            "PATCH",
			route:             "/v1/user/",
			body:              nil,
			authReq:           false,
			cred:              nil,
			expectedError:     true,
			expectedErrorDesc: "You must be logged in to access this resource.",
			expectedCode:      http.StatusUnauthorized,
		},
		{
			// request invalid body field
			method:            "PATCH",
			route:             "/v1/user/",
			body:              []byte(`{}`),
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusBadRequest,
		},
		{
			// request invalid body field
			method:            "PATCH",
			route:             "/v1/user/",
			body:              &model.ProfileUpdateRequest{Name: "murat", Email: "murat"},
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     true,
			expectedErrorDesc: "",
			expectedCode:      http.StatusBadRequest,
		},
		{
			// successful
			method:            "PATCH",
			route:             "/v1/user/",
			body:              &model.ProfileUpdateRequest{Name: "murat", Email: "murat@gmail.com"},
			authReq:           true,
			cred:              &model.LoginRequest{Email: "johndoe@gmail.com", Password: "12345678"},
			expectedError:     false,
			expectedErrorDesc: "",
			expectedCode:      http.StatusOK,
		},
	}

	for _, test := range tests {
		c, resp := makeRequest(test.method, test.route, test.body, test.authReq, test.cred)

		if test.expectedError && test.expectedErrorDesc != "" {
			err := AuthenticateUser(srv.handleUpdateProfile)(c)
			assert.NotNil(t, err)

			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)

			assert.Equal(t, test.expectedCode, he.Code)
			assert.Equal(t, test.expectedErrorDesc, he.Message)

			continue

		}

		if assert.NoError(t, AuthenticateUser(srv.handleUpdateProfile)(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}
	}
}

func TestHandleDeleteProfile(t *testing.T) {
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
			method:            "DELETE",
			route:             "/v1/user/",
			body:              nil,
			authReq:           false,
			cred:              nil,
			expectedError:     true,
			expectedErrorDesc: "You must be logged in to access this resource.",
			expectedCode:      http.StatusUnauthorized,
		},
		{
			// successful
			method:            "DELETE",
			route:             "/v1/user/",
			body:              nil,
			authReq:           true,
			cred:              &model.LoginRequest{Email: "murat@gmail.com", Password: "12345678"},
			expectedError:     false,
			expectedErrorDesc: "",
			expectedCode:      http.StatusSeeOther,
		},
	}

	for _, test := range tests {
		c, resp := makeRequest(test.method, test.route, test.body, test.authReq, test.cred)

		if test.expectedError && test.expectedErrorDesc != "" {
			err := AuthenticateUser(srv.handleDeleteProfile)(c)
			assert.NotNil(t, err)

			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)

			assert.Equal(t, test.expectedCode, he.Code)
			assert.Equal(t, test.expectedErrorDesc, he.Message)

			continue

		}

		if assert.NoError(t, AuthenticateUser(srv.handleDeleteProfile)(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}
	}

	// recreate user to not affect other tests
	c, resp := makeRequest("POST", "/v1/auth/register", model.RegisterRequest{
		Name:            "john",
		Email:           "johndoe@gmail.com",
		Password:        "12345678",
		PasswordConfirm: "12345678",
	}, false, nil)

	if assert.NoError(t, srv.handleRegister(c)) {
		require.Equal(t, http.StatusCreated, resp.Code)
	}
}
