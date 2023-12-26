package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ProtectUser(c *fiber.Ctx) error{
	usernameFromContext := c.UserContext().Value("username").(string)
	usernameFromURL     := c.Params("username")

	fmt.Println("username url:", usernameFromURL, "path:", c.Path(), "all params:", c.AllParams())

	if usernameFromContext != usernameFromURL{
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"err": fmt.Sprintf("User %s is not allowed to "+
			"view information from user %s", usernameFromContext, usernameFromURL)})
	}

	return c.Next()
}