package server

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/orhanfatih/blog-api/model"
	"github.com/stretchr/testify/assert"
)

func TestHandleRegister(t *testing.T) {

	tests := []struct {
		method        string
		route         string
		body          interface{}
		authReq       bool
		cred          *model.LoginRequest
		expectedError bool
		expectedCode  int
	}{
		{
			// request invalid body field
			method:        "POST",
			route:         "/v1/auth/register",
			body:          []byte(`{}`),
			authReq:       false,
			cred:          nil,
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
		{
			// request body not filled fully
			method: "POST",
			route:  "/v1/auth/register",
			body: model.RegisterRequest{
				Name:     "john",
				Email:    "johndoe@gmail.com",
				Password: "12345678",
				// PasswordConfirm: "12345698",
			},
			authReq:       false,
			cred:          nil,
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
		{
			method: "POST",
			route:  "/v1/auth/register",
			body: model.RegisterRequest{
				Name: "john",
				// Email:    "johndoe@gmail.com",
				Password:        "12345678",
				PasswordConfirm: "12345678",
			},
			authReq:       false,
			cred:          nil,
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
		{
			// field invalid values
			method: "POST",
			route:  "/v1/auth/register",
			body: model.RegisterRequest{
				Name:            "john",
				Email:           "c",
				Password:        "12345678",
				PasswordConfirm: "12345678",
			},
			authReq:       false,
			cred:          nil,
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
		{
			// successful
			method: "POST",
			route:  "/v1/auth/register",
			body: model.RegisterRequest{
				Name:            "john",
				Email:           "johndoe@gmail.com",
				Password:        "12345678",
				PasswordConfirm: "12345678",
			},
			authReq:       false,
			cred:          nil,
			expectedError: false,
			expectedCode:  http.StatusCreated,
		},
		{
			// duplicate unique constraint
			method: "POST",
			route:  "/v1/auth/register",
			body: model.RegisterRequest{
				Name:            "john",
				Email:           "johndoe@gmail.com",
				Password:        "12345678",
				PasswordConfirm: "12345678",
			},
			authReq:       false,
			cred:          nil,
			expectedError: false,
			expectedCode:  http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		c, resp := makeRequest(test.method, test.route, test.body, test.authReq, test.cred)

		if assert.NoError(t, srv.handleRegister(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}
	}
}

func TestHandleLogin(t *testing.T) {

	tests := []struct {
		method        string
		route         string
		body          interface{}
		authReq       bool
		cred          *model.LoginRequest
		expectedError bool
		expectedCode  int
	}{
		{
			// request invalid body field
			method:        "POST",
			route:         "/v1/auth/login",
			body:          []byte(`{}`),
			authReq:       false,
			cred:          nil,
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
		{
			// missing body fields
			method: "POST",
			route:  "/v1/auth/login",
			body: model.LoginRequest{
				Password: "12345678",
			},
			authReq:       false,
			cred:          nil,
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
		{
			// missing body fields
			method: "POST",
			route:  "/v1/auth/login",
			body: model.LoginRequest{
				Email: "johndoe@gmail.com",
			},
			authReq:       false,
			cred:          nil,
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
		{
			// invalid body field values
			method: "POST",
			route:  "/v1/auth/login",
			body: model.LoginRequest{
				Email:    "213231",
				Password: "1234",
			},
			authReq:       false,
			cred:          nil,
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
		{
			// invalid body field values
			method: "POST",
			route:  "/v1/auth/login",
			body: model.LoginRequest{
				Email:    "johndoe@gmail.com",
				Password: "1234",
			},
			authReq:       false,
			cred:          nil,
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
		},
		{
			// invalid credientials
			method: "POST",
			route:  "/v1/auth/login",
			body: model.LoginRequest{
				Email:    "johndoe@gmail.com",
				Password: "87654321",
			},
			authReq:       false,
			cred:          nil,
			expectedError: false,
			expectedCode:  http.StatusBadRequest,
		},
		{
			// successful
			method: "POST",
			route:  "/v1/auth/login",
			body: model.LoginRequest{
				Email:    "johndoe@gmail.com",
				Password: "12345678",
			},
			authReq:       false,
			cred:          nil,
			expectedError: false,
			expectedCode:  http.StatusOK,
		},
	}

	for _, test := range tests {
		c, resp := makeRequest(test.method, test.route, test.body, test.authReq, test.cred)

		if assert.NoError(t, srv.handleLogin(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}
	}
}

func TestHandleLogout(t *testing.T) {
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
			route:             "/v1/auth/logout",
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
			route:             "/v1/auth/logout",
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

		if test.expectedError && test.expectedErrorDesc != "" {
			err := AuthenticateUser(srv.handleLogout)(c)
			assert.NotNil(t, err)

			he, ok := err.(*echo.HTTPError)
			assert.True(t, ok)

			assert.Equal(t, test.expectedCode, he.Code)
			assert.Equal(t, test.expectedErrorDesc, he.Message)

			continue

		}

		if assert.NoError(t, AuthenticateUser(srv.handleLogout)(c)) {
			assert.Equal(t, test.expectedCode, resp.Code)
		}
	}
}
