package server

import (
	"log"

	"github.com/labstack/echo"
)

func RespondWithError(c echo.Context, code int, err string) error {
	log.Printf("[http] error: %s %s: %s", c.Request().Method, c.Request().URL.Path, err)
	return RespondWithJSON(c, code, map[string]string{"Status": "Error", "Message": err})
}

func RespondWithJSON(c echo.Context, code int, payload interface{}) error {
	return c.JSON(code, payload)
}
