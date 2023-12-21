package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ProtectSignin(c *fiber.Ctx) error{
	fmt.Println("path", c.Path())
	fmt.Println(strings.HasSuffix(c.Path(), "/signin"))
	
	return c.Next()
}