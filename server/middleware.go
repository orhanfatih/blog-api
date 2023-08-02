package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

func CreateToken(id uint, expiration time.Duration, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
		"iss": "blog-api",
		"iat": time.Now().UTC().Unix(),
		"exp": time.Now().Add(expiration).UTC().Unix(),
	}

	// create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method.
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid signing method: %v", token.Method)
		}

		secret := []byte(os.Getenv("JWT_SECRET"))
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the access token from the cookie
		cookie, err := c.Cookie("access-token")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// Validate the token
		claims, err := ValidateToken(cookie.Value)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// Extract the user ID from the claims
		userID, err := strconv.Atoi(fmt.Sprint(claims["sub"]))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// Set the user ID in the request context for later use in the route handler
		c.Set("userID", userID)

		// Proceed to the next middleware or the route handler
		return next(c)
	}
}
