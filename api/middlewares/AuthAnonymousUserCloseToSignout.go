package middlewares

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"MessengerV2/api/database"
)

//Function to read the
func AuthAnonymousUserCloseToSignout(c *fiber.Ctx) error {
	expiry, _ := c.UserContext().Value("token").(jwt.MapClaims)["exp"].(int64)
	isAnonymous, _ := c.UserContext().Value("token").(jwt.MapClaims)["is_anonymous"].(bool)
	username, _ := c.UserContext().Value("token").(jwt.MapClaims)["username"].(string)

	//If the anonymous user's cookie has 10 minutes (600 seconds) or less before expiring, remove them 
	//from the database, and expire the cookie.
	if expiry - time.Now().Unix() <= 600 && isAnonymous {
		if err := database.DeleteAnonymousUser(username); err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		c.Cookie(&fiber.Cookie{
			MaxAge: 0,
			Value:  "",
			Name:   "user",
		})

		return c.Status(http.StatusUnauthorized).SendString("A valid cookie is required for entry.")
	} 

	return c.Next()
}