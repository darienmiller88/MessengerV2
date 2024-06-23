package middlewares

import (
	"fmt"
	"net/http"

	"MessengerV2/api/database"
	"MessengerV2/api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

//Check if a user is in the group chat they are trying to make changes to. This middleware prevents users from
//making calls to chats they don't belong in.
func AuthCheckUser(c *fiber.Ctx) error{
	username, usernameErr := c.UserContext().Value("token").(jwt.MapClaims)["username"].(string)

	if !usernameErr {
		return c.Status(http.StatusUnprocessableEntity).SendString("Could not parse \"username\" field.")
	}
	
	id := c.Params("chatid")

	userChat := models.UserChat{}
	if err := database.GetDB().Get(&userChat, "SELECT * FROM user_chats WHERE username=$1 AND chat_id=$2", username, id); err != nil{
		fmt.Println("err auth middleare:", err)
		return c.Status(http.StatusForbidden).SendString(err.Error())
	}

	return c.Next()
}