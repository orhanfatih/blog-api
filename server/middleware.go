package server

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(id uint, expiration time.Duration, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
		"iss": "blog-api",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(expiration).Unix(),
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
