package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ProtectSignin(c *fiber.Ctx) error{
	_, err := retrieveTokenFromCookie(c.Cookies("user"))
	pageHit := strings.Contains(c.Path(), "/signin") || strings.Contains(c.Path(), "/signup")

	//If the user is currently signed in, and they try to access the signin or signup route, block the request
	if err == nil && pageHit {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"err": fmt.Sprintf("Cannot access %s while signed in.", c.Path())})
	}
	
	return c.Next()
}