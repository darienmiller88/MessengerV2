package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"MessengerV2/api/database"
	"MessengerV2/api/models"
)

func AuthCheckUser(c *fiber.Ctx) error{
	username, usernameErr := c.UserContext().Value("token").(jwt.MapClaims)["username"].(string)

	if !usernameErr {
		return c.Status(http.StatusUnprocessableEntity).SendString("Could not parse \"username\" field.")
	}
	
	id := c.Params("id")

	userChat := models.UserChat{}
	if err := database.GetDB().Select(&userChat, "SELECT * FROM user_chats WHERE username=$1 AND chat_id=$2", username, id); err != nil{
		return c.Status(http.StatusForbidden).SendString(err.Error())
	}

	return c.Next()
}