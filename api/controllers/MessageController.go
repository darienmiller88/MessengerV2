package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/pusher/pusher-http-go/v5"

	"MessengerV2/api/database"
	"MessengerV2/api/models"
	"MessengerV2/api/pusherclient"
	"MessengerV2/api/cloudinary"
)

const MAX_SIZE int64 = 5 * 1024 * 1024 //Max number of bytes, which is 5mb or 5,248,000 bytes
const public string = "Public"

type MessageController struct{
	pusherClient pusher.Client
	db           *sqlx.DB
}

func (m *MessageController) Init(){
	m.pusherClient = pusherclient.GetPusherClient()
	m.db           = database.GetDB()
}

func (m *MessageController) UploadImageAsMessage(c *fiber.Ctx) error{
	chatId    := c.Params("chat-id", public)
	file, err := c.FormFile("file")

	if err != nil{
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if file.Size > MAX_SIZE{
		return c.Status(http.StatusBadRequest).SendString("File too big.")
	}

	res, err := cloudinary.UploadImage(file)
	
	if err != nil{
		return c.SendString(err.Error())
	}

	message := models.Message{DB: m.db}

	message.InitCreatedAt()
	message.Username        = c.FormValue("username")
	message.DisplayName     = c.FormValue("display_name")
	message.MessageContent  = c.FormValue("message_content")
	message.MessageDate     = time.Now().Format("2006-01-02 3:4:5 pm")
	message.ImageURL.Valid  = res != ""
	message.ImageURL.String = res

	if err := message.Validate(); err != nil{
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	m.pusherClient.Trigger(chatId, "message", message)
	
	message, _ = database.InsertMessage(message)

	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) PostMessage(c *fiber.Ctx) error{
	message := models.Message{DB: m.db}
	chatId  := c.Params("chat-id", public)

	if err := c.BodyParser(&message); err != nil{
		fmt.Println("err parsing message:", err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if err := message.Validate(); err != nil{
		fmt.Println("err validating message:", err)
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	
	message.InitCreatedAt()
	message.MessageDate = time.Now().Format("2006-01-02 3:4:5 pm")
	
	if err := m.pusherClient.Trigger(chatId, "message", message); err != nil{
		fmt.Println("err broadcasting messages:", err)
	}

	message, err := database.InsertMessage(message)

	if err != nil{
		fmt.Println("err inserting message:", err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) GetMessageByID(c *fiber.Ctx) error{
	id      := c.Params("id")
	message := models.Message{}

	if err := m.db.Get(&message, "SELECT * FROM messages WHERE id=$1", id); err != nil{
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"err": fmt.Sprintf("Message with id %s does not exist.", id)})
	}

	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) GetMessageHistory(c *fiber.Ctx) error{
	messages              := []models.Message{}
	username, usernameErr := c.UserContext().Value("token").(jwt.MapClaims)["username"].(string) 

	if !usernameErr{
		return c.Status(http.StatusBadRequest).SendString("Could not parse field \"username\" in token.")
	}

	if err := m.db.Select(&messages, "SELECT * FROM messages WHERE username=$1", username); err != nil{
		return c.Status(http.StatusNotFound).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(messages)
}

func (m *MessageController) GetGroupChatMessages(c *fiber.Ctx) error{
	id       := c.Params("id")
	messages := []models.Message{}

	if err := m.db.Select(&messages, "SELECT * FROM messages WHERE chat_id=$1 ORDER BY id ASC", id); err != nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(messages)
}

func (m *MessageController) GetPublicMessages(c *fiber.Ctx) error{
	messages := []models.Message{}

	// SELECT messages.*, users.profile_picture FROM messages JOIN users ON messages.username = users.username;
	if err := m.db.Select(&messages, "SELECT * FROM messages WHERE receiver IS NULL AND chat_id IS NULL ORDER BY id ASC"); err != nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(messages)
}

func (m *MessageController) DeleteMessage(c *fiber.Ctx) error{
	id      := c.Params("id", public)
	message := models.Message{}

	if err := c.BodyParser(&message); err != nil{
		c.Status(http.StatusInternalServerError).JSON(err)
	}

	if err := m.pusherClient.Trigger("public", "delete_message", message); err != nil{
		fmt.Println("err broadcasting messages:", err)
	}

	result, err     := m.db.Exec("DELETE FROM messages WHERE id=$1 and username=$2", id, message.Username)
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 || err != nil{
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(message)
}

//Function to allow clients on the front end to know when someone is typing.
func (m *MessageController) UserTyping(c *fiber.Ctx) error{
	chatId := c.Params("chat-id", public)
	data   := struct{
		Username string `json:"username"`
	}{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := m.pusherClient.Trigger(chatId, "user_typing", data.Username); err != nil{
		fmt.Println("err broadcasting messages:", err)
	}

	return c.Status(http.StatusOK).JSON(data)
}