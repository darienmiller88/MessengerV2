package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)
func Auth(c *fiber.Ctx) error{
	token, err := retrieveTokenFromCookie(c.Cookies("user"))//Retrieve value from cookie and parse token from it.

	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString(err.Error())
	}

	//Parse the token and pull out the username from the token.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.SetUserContext(context.WithValue(c.Context(), "username", claims["username"]))
	}

	return c.Next()
}

func retrieveTokenFromCookie(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return token, nil
}