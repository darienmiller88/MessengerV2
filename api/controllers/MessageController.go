package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pusher/pusher-http-go/v5"

	"MessengerV2/api/database"
	"MessengerV2/api/models"
	"MessengerV2/api/pusherclient"
)

type MessageController struct{
	pusherClient pusher.Client
	db           *sqlx.DB
}

func (m *MessageController) Init(){
	m.pusherClient = pusherclient.GetPusherClient()
	m.db           = database.GetDB()
}

func (m *MessageController) PostMessage(c *fiber.Ctx) error{
	message := models.Message{DB: m.db}

	if err := c.BodyParser(&message); err != nil{
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	// if err := message.Validate(); err != nil{
	// 	return c.Status(http.StatusBadRequest).JSON(err)
	// }
	
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	message.MessageDate = time.Now().Format("2006-01-02 3:4:5 pm")
	// tx := m.db.MustBegin()
	
	result, err := m.db.NamedExec("INSERT INTO messages (message_content, message_date, created_at, updated_at, username) " +
	"VALUES(:message_content, :message_date, :created_at, :updated_at, :username)", &message)
	
	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}
	
	fmt.Println("result", result)
	fmt.Println("message:", message)
	// tx.Commit()
	// if err := m.pusherClient.Trigger("public", "public_message", message); err != nil{
	// 	fmt.Println("err broadcasting messages:", err)
	// }

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
	username := c.Params("username")
	message  := models.Message{}

	if err := m.db.Get(&message, "SELECT * FROM messages WHERE username=$1", username); err != nil{
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"err": fmt.Sprintf("Username %s does not exist.", username)})
	}

	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) GetMessages(c *fiber.Ctx) error{
	messages := []models.Message{}

	if err := m.db.Select(&messages, "SELECT * FROM messages"); err != nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(messages)
}

func (m *MessageController) DeleteMessage(c *fiber.Ctx) error{
	id      := c.Params("id")
	message := models.Message{}

	if err := c.BodyParser(&message); err != nil{
		c.Status(http.StatusInternalServerError).JSON(err)
	}

	fmt.Println("id:", id)
	
	err := m.pusherClient.Trigger("public", "delete_public_message", message)

	if err != nil{
		fmt.Println("err broadcasting messages:", err)
	}

	return c.Status(http.StatusOK).JSON(message)
}

func (m *MessageController) UserTyping(c *fiber.Ctx) error{
	data := make(map[string]interface{})

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	err := m.pusherClient.Trigger("public", "user_typing", data["username"])

	if err != nil{
		fmt.Println("err broadcasting messages:", err)
	}

	return c.Status(http.StatusOK).JSON(data)
}