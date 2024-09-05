package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

//Prevent users from being able to view the information of another user when using "username" as a route parameter.
func ProtectUser(c *fiber.Ctx) error{
	usernameFromContext := c.UserContext().Value("token").(jwt.MapClaims)["username"]
	usernameFromURL     := c.Params("username")

	if usernameFromContext != usernameFromURL{
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"err": fmt.Sprintf("User %s is not allowed to "+
			"view information from user %s", usernameFromContext, usernameFromURL)})
	}

	return c.Next()
}